package output

import (
	"github.com/Design-Pattern-Go-Implementation/monitor/config"
	"github.com/Design-Pattern-Go-Implementation/monitor/plugin"
	"reflect"
)

// Plugin 抽象的输出接口
type Plugin interface {
	plugin.Plugin
	Output(event *plugin.Event) error
}

// NewPlugin 输出插件工厂方法
func NewPlugin(config config.Output) (Plugin, error) {
	// output这个包下面的map
	outputType, ok := Type[config.PluginType]
	if !ok {
		return nil, plugin.ErrUnknownPlugin
	}

	outputPlugin := reflect.New(outputType)
	ctx := reflect.ValueOf(config.Ctx)
	outputPlugin.MethodByName("SetContext").Call([]reflect.Value{ctx})
	return outputPlugin.Interface().(Plugin), nil
}
