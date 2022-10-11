package proxy

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"strings"
	"text/template"
)

// 生成代理类的文件模板
const proxyTpl = `package {{.Package}}
type {{ .ProxyStructName }}Proxy struct {
	child *{{ .ProxyStructName }}
}
func New{{ .ProxyStructName }}Proxy(child *{{ .ProxyStructName }}) *{{ .ProxyStructName }}Proxy{
	return &{{ .ProxyStructName }}Proxy{child: child}
}
{{ range .Methods }}
func (p *{{$.ProxyStructName}}Proxy) {{ .Name }}({{ .Params }}) ({{ .Results }}) {
	// before 这里可能会有一些统计的逻辑
	start := time.Now()
	{{ .ResultNames }} = p.child.{{ .Name }}({{ .ParamNames }})
	// after 这里可能也有一些监控统计的逻辑
	log.Printf("user login cost time: %s", time.Now().Sub(start))
	return {{ .ResultNames }}
}
{{ end }}
`

/*
关键点：1.给谁代理？ 2.代理它的哪个方法？

*/
type proxyData struct {
	// 包名
	Package string
	// 需要代理的类名
	ProxyStructName string
	// 需要代理的方法
	Methods []*proxyMethod
}

type proxyMethod struct {
	// 被代理的方法名字在程序中需要被调用
	Name string
	// 被代理的方法的参数类型
	Params string
	// 被代理的方法的参数名
	ParamNames string
	// 被代理的方法的返回值类型
	Results string
	// 返回值名字
	ResultNames string
}

func generate(file string) (string, error) {
	fset := token.NewFileSet()
	// 手动解析.go文件程序，获取其中定义的常量的值和之后的注释
	// f存储着解析好的接口，
	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		return "", err
	}
	// 获取代理需要的信息
	data := proxyData{
		Package: f.Name.Name,
	}
	// 构建注释和 node的关系
	cmap := ast.NewCommentMap(fset, f, f.Comments)
	for node, group := range cmap {
		//从注释 @proxy 接口名，获取接口名称
		name := getProxyInterfaceName(group)
		if name == "" {
			continue
		}
		// 获取代理的类名
		data.ProxyStructName = node.(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Name.Name
		// 从文件中查找接口
		obj := f.Scope.Lookup(name)
		// 类型转化
		// 类型转换，注意: 这里没有对断言进行判断，可能会导致 panic
		t := obj.Decl.(*ast.TypeSpec).Type.(*ast.InterfaceType)
		// 获取到接口的方法列表
		for _, field := range t.Methods.List {
			fc := field.Type.(*ast.FuncType)
			// 代理的方法
			method := &proxyMethod{
				Name: field.Names[0].Name,
			}
			method.Params, method.ParamNames = getParamsOrResults(fc.Params)
			method.Results, method.ResultNames = getParamsOrResults(fc.Results)
			data.Methods = append(data.Methods, method)

		}

	}

	// 根据proxyTpl 模板生成文件
	tpl, err := template.New("").Parse(proxyTpl)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, data); err != nil {
		return "", err
	}
	// 使用 go fmt 对生成的代码进行格式化
	src, err := format.Source(buf.Bytes())

	if err != nil {
		return "", err
	}
	return string(src), nil
}

func getProxyInterfaceName(groups []*ast.CommentGroup) string {
	for _, commentGroup := range groups {
		for _, comment := range commentGroup.List {
			if strings.Contains(comment.Text, "@proxy") {
				interfaceName := strings.TrimLeft(comment.Text, "// @proxy ")
				return strings.TrimSpace(interfaceName)
			}
		}
	}
	return ""
}

// getParamsOrResults 获取参数或者是返回值
// 返回带类型的参数，以及不带类型的参数，以逗号间隔
func getParamsOrResults(fields *ast.FieldList) (string, string) {
	var (
		params     []string
		paramNames []string
	)

	for i, param := range fields.List {
		// 循环获取所有的参数名
		var names []string
		for _, name := range param.Names {
			names = append(names, name.Name)
		}

		if len(names) == 0 {
			names = append(names, fmt.Sprintf("r%d", i))
		}

		paramNames = append(paramNames, names...)

		// 参数名加参数类型组成完整的参数
		param := fmt.Sprintf("%s %s",
			strings.Join(names, ","),
			param.Type.(*ast.Ident).Name,
		)
		params = append(params, strings.TrimSpace(param))
	}

	return strings.Join(params, ","), strings.Join(paramNames, ",")
}
