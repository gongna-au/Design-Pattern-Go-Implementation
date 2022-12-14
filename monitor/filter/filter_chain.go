package filter

import (
	"github.com/Design-Pattern-Go-Implementation/monitor/plugin"
)

/*
责任链模式
*/

// Chain Filter链，按顺序调用
// 具体的插件 实现了Plugin 接口
type Chain struct {
	filters []Plugin
}

func NewChain(filters []Plugin) *Chain {
	return &Chain{filters: filters}
}

func (c *Chain) Filter(event *plugin.Event) *plugin.Event {

	for _, filter := range c.filters {
		event = filter.Filter(event)
	}
	return event
}

func (c *Chain) Install() {
	for _, filter := range c.filters {
		filter.Install()
	}
}

func (c *Chain) Uninstall() {
	for _, filter := range c.filters {
		filter.Uninstall()
	}
}

func (c *Chain) SetContext(ctx plugin.Context) {
}
