package creation

// 工厂方法
// 当对象的创建逻辑比较复杂，
// 不只是简单的 new 一下就可以，
// 而是要组合其他类对象，做各种初始化操作的时候，推荐使用工厂方法模式，将复杂的创建逻辑拆分到多个工厂类中，让每个工厂类都不至于过于复杂

// IRuleConfigParserFactory 工厂方法接口
type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

// yamlRuleConfigParserFactory yamlRuleConfigParser 的工厂类
type yamlRuleConfigParserFactory struct {
}

// CreateParser CreateParser
func (y yamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return yamlRuleConfigParser{}
}

// jsonRuleConfigParserFactory jsonRuleConfigParser 的工厂类
type jsonRuleConfigParserFactory struct {
}

// CreateParser CreateParser
func (j jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

// NewIRuleConfigParserFactory 用一个简单工厂封装工厂方法
func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return jsonRuleConfigParserFactory{}
	case "yaml":
		return yamlRuleConfigParserFactory{}
	}
	return nil
}

// 还可是这样封装

/* var IRuleConfigParserFactoryMap map[string]reflect.Type

func init() {
	IRuleConfigParserFactoryMap = make(map[string]reflect.Type)
	IRuleConfigParserFactoryMap["json"] = reflect.TypeOf(jsonRuleConfigParserFactory{})
	IRuleConfigParserFactoryMap["yaml"] = reflect.TypeOf(yamlRuleConfigParserFactory{})
}

func NewIRuleConfigParserFactory(t string) (IRuleConfigParserFactory, error) {
	v, ok := IRuleConfigParserFactoryMap[t]
	if ok {
		result := reflect.ValueOf(v)
		return result.Interface().(IRuleConfigParserFactory), nil
	} else {
		return nil, errors.New("IRuleConfigParserFactory " + t + "is not exist")
	}
} */

/*
type IRuleConfigParser interface {
	Parse(data []byte)
}

// jsonRuleConfigParser jsonRuleConfigParser
type jsonRuleConfigParser struct {
}

// Parse Parse
func (J jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// yamlRuleConfigParser yamlRuleConfigParser
type yamlRuleConfigParser struct {
}

// Parse Parse
func (Y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}
*/
