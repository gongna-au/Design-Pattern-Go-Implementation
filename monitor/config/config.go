package config

import (
	"github.com/Design-Pattern-Go-Implementation/monitor/plugin"
)

type Type uint8

const (
	JsonType Type = iota
	YamlType
)

// 所有的插件都要有个属性标明自己的类型 PluginType 是 int 类型
// 用Name标志插件的名字
// 每个插件需要从配置文件中加载配置，可以把这个实现放在具体的插件实现
// 但是不同的插件加载配置的方式都是一样的，所以可以放在元插件item  去实现（或者放在属性里面）

type item struct {
	Name       string         `json:"name" yaml:"name"`
	PluginType string         `json:"type" yaml:"type"`
	Ctx        plugin.Context `json:"context" yaml:"context"`
	loadConf   func(conf string, item interface{}) error
}

type Input item

func (i *Input) Load(conf string) error {
	return i.loadConf(conf, i)
}

type Filter item

func (f *Filter) Load(conf string) error {
	return f.loadConf(conf, f)
}

type Output item

func (o *Output) Load(conf string) error {
	return o.loadConf(conf, o)
}

type Pipeline struct {
	// item 只是为了继承了loadConf（）
	item    `yaml:",inline"` // yaml嵌套时需要加上,inline
	Input   Input            `json:"input" yaml:"input"`
	Filters []Filter         `json:"filters" yaml:"filters,flow"`
	Output  Output           `json:"output" yaml:"output"`
}

func (p *Pipeline) Load(conf string) error {
	return p.loadConf(conf, p)
}
