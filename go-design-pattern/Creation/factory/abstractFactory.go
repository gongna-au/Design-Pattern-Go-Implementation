package factory

// 抽象工厂

// IRuleConfigParser IRuleConfigParser
type IRulesConfigParser interface {
	Parse(data []byte)
}

// jsonRuleConfigParser jsonRuleConfigParser
type jsonRulesConfigParser struct{}

// Parse Parse
func (j jsonRulesConfigParser) Parse(data []byte) {
	panic("implement me")
}

// ISystemConfigParser ISystemConfigParser
type ISystemConfigParser interface {
	ParseSystem(data []byte)
}

// jsonSystemConfigParser jsonSystemConfigParser
type jsonSystemConfigParser struct{}

// Parse Parse
func (j jsonSystemConfigParser) ParseSystem(data []byte) {
	panic("implement me")
}

// IConfigParserFactory 工厂方法接口
type IConfigParserFactory interface {
	CreateRuleParser() IRuleConfigParser
	CreateSystemParser() ISystemConfigParser
}

type jsonConfigParserFactory struct{}

func (j jsonConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return jsonRulesConfigParser{}
}

func (j jsonConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return jsonSystemConfigParser{}
}

// 注意观察区别，一个的本质是 RuleConfigParser 一个是SystemConfigParser
// 站在客户端视角:
// 工厂方法是：客户端调用一个NewFactory 函数得到不同的工厂类，不同的工厂类就可以直接生成具体的产品，产品的功能都是一样的
// 抽象工厂是：客户端拿着一个抽象的工厂类，调用不同的函数得到不同的产品
func NewJsonConfigParserFactory() jsonConfigParserFactory {
	return jsonConfigParserFactory{}
}
