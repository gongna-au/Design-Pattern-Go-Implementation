package pipeline

import (
	"github.com/Design-Pattern-Go-Implementation/monitor/plugin"
	"reflect"
)

// plugin.Types  map[string]reflect.Type
// Type input插件类型
// 如果想要在map里面存储 不同类型的实例，那么map的类型一定是下面这种 map[]reflect.Type
// 存储插件类型的map 并且在程序运行的时候就已经被初始化
var Type = make(plugin.Types)

func init() {
	Type["simple"] = reflect.TypeOf(SimplePipeline{})
	Type["pool"] = reflect.TypeOf(PoolPipeline{})
}
