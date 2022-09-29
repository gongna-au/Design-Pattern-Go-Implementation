package input

import (
	"github.com/Design-Pattern-Go-Implementation/monitor/config"
	"github.com/Design-Pattern-Go-Implementation/monitor/plugin"
	"reflect"
)

/*
策略模式
*/
// Plugin 抽象的输入接口
// 具体的过滤插件去实现 Install()  Uninstall() SetContext(ctx Context) 以及Input()
// 然后把具体的插件存储在一个map中
// 由工厂方法来返回我们需要的插件
// 然后 NewPlugin 返回
type Plugin interface {
	plugin.Plugin
	Input() (*plugin.Event, error)
}

// 工厂方法 就只负责返回我们需要的插件
// 定义产生具体插件的方式
// NewPlugin 输入插件工厂方法
// map[string]reflect.Type
// 一般的map存储同一种类型实例
// reflect.Typeof()转化为统一的类型
// 在需要的时候，转化为具体的实例
func NewPlugin(config config.Input) (Plugin, error) {
	//得到插件类型
	inputType, ok := Type[config.PluginType]
	if !ok {
		return nil, plugin.ErrUnknownPlugin
	}
	//根据类型转化为具体的插件
	inputPlugin := reflect.New(inputType)
	ctx := reflect.ValueOf(config.Ctx)
	inputPlugin.MethodByName("SetContext").Call([]reflect.Value{ctx})
	// inputPlugin的得到的具体的类型，但是是reflect.Value
	// 可以调用Interface()得到实例interface 再把interface 转化为实现的接口便于赋值
	return inputPlugin.Interface().(Plugin), nil
}
