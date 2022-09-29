package filter

import (
	"github.com/Design-Pattern-Go-Implementation/monitor/config"
	"github.com/Design-Pattern-Go-Implementation/monitor/plugin"
	"reflect"
)

// Plugin 抽象的过滤接口
// 具体的过滤插件去实现 Install()  Uninstall() SetContext(ctx Context) 以及Filter(event *plugin.Event)
type Plugin interface {
	plugin.Plugin
	Filter(event *plugin.Event) *plugin.Event
}

// NewPlugin 过滤插件工厂方法
func NewPlugin(config config.Filter) (Plugin, error) {
	filterType, ok := Type[config.PluginType]
	if !ok {
		return nil, plugin.ErrUnknownPlugin
	}
	filterPlugin := reflect.New(filterType)
	ctx := reflect.ValueOf(config.Ctx)
	filterPlugin.MethodByName("SetContext").Call([]reflect.Value{ctx})
	return filterPlugin.Interface().(Plugin), nil
}
