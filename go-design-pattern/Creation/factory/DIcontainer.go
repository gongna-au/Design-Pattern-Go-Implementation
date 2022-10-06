package factory

import (
	"fmt"
	"reflect"
)

// DI容器：
// 把有依赖关系的类放到容器中，解析出这些类的实例，就是依赖注入。
// Container DI 容器，，随着业务的扩展，我们如果把所有实例都在main函数里生成，main函数将变得越来越臃肿。
// 而基础服务的实例，如果在其他包里需要引入，你就得给每个需要用到服务的地方，通过参数的方式传递这些实例
// DI让main函数变得优雅、 可以管理全局服务实例、减少传递如config这样的基础实例
// 创建容器、往容器中添加一些某个实例初始化的方法
type Container struct {
	// 假设一种类型只能有一个 provider 提供
	providers map[reflect.Type]provider

	// 缓存以生成的对象
	results map[reflect.Type]reflect.Value
}

// Provide 对象提供者，需要传入一个对象的工厂方法，后续会用于对象的创建
func (c *Container) Provide(constructor interface{}) error {
	v := reflect.ValueOf(constructor)
	if v.Kind() != reflect.Func {
		return fmt.Errorf("constructor must be a func")
	}
	vt := v.Type()
	// 获取参数
	params := make([]reflect.Type, vt.NumIn())
	for i := 0; i < vt.NumIn(); i++ {
		params[i] = vt.In(i)
	}
	// 获取返回值
	results := make([]reflect.Type, vt.NumOut())
	for i := 0; i < vt.NumOut(); i++ {
		results[i] = vt.Out(i)
	}
	// provide 存储类型的参数
	provider := provider{
		value:  v,
		params: params,
	}
	// 保存不同类型的 provider
	// 每次往容器中添加一个函数，用到前面的结果
	for _, result := range results {
		// 判断返回值是不是 error
		if isError(result) {
			continue
		}

		if _, ok := c.providers[result]; ok {
			return fmt.Errorf("%s had a provider", result)
		}

		c.providers[result] = provider
	}
	return nil

}

// Invoke 函数执行入口
func (c *Container) Invoke(function interface{}) error {
	v := reflect.ValueOf(function)

	// 仅支持函数 provider
	if v.Kind() != reflect.Func {
		return fmt.Errorf("constructor must be a func")
	}

	vt := v.Type()

	// 获取参数
	var err error
	params := make([]reflect.Value, vt.NumIn())
	for i := 0; i < vt.NumIn(); i++ {
		params[i], err = c.buildParam(vt.In(i))
		if err != nil {
			return err
		}
	}

	v.Call(params)

	// 获取 providers
	return nil
}

// buildParam 构建参数
// 1. 从容器中获取 provider
// 2. 递归获取 provider 的参数值
// 3. 获取到参数之后执行函数
// 4. 将结果缓存并且返回结果
func (c *Container) buildParam(param reflect.Type) (val reflect.Value, err error) {
	if result, ok := c.results[param]; ok {
		return result, nil
	}
	//
	provider, ok := c.providers[param]
	if !ok {
		return reflect.Value{}, fmt.Errorf("can not found provider: %s", param)
	}

	params := make([]reflect.Value, len(provider.params))
	for i, p := range provider.params {
		params[i], err = c.buildParam(p)
	}

	results := provider.value.Call(params)
	for _, result := range results {
		// 判断是否报错
		if isError(result.Type()) && !result.IsNil() {
			return reflect.Value{}, fmt.Errorf("%s call err: %+v", provider, result)
		}

		if !isError(result.Type()) && !result.IsNil() {
			c.results[result.Type()] = result
		}
	}
	return c.results[param], nil
}

// isError 判断是否是 error 类型
func isError(t reflect.Type) bool {
	if t.Kind() != reflect.Interface {
		return false
	}
	return t.Implements(reflect.TypeOf(reflect.TypeOf((*error)(nil)).Elem()))
}

type provider struct {
	value reflect.Value

	params []reflect.Type
}

// New 创建一个容器
func New() *Container {
	return &Container{
		providers: map[reflect.Type]provider{},
		results:   map[reflect.Type]reflect.Value{},
	}
}
