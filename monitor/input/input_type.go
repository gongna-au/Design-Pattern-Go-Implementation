package input

import (
	"github.com/Design-Pattern-Go-Implementation/monitor/plugin"
	"reflect"
)

// Type input插件类型
// 为每个抽象的插件工厂创建了一个map 实例来存储具体的实例
var Type = make(plugin.Types)

func init() {
	Type["memory_mq"] = reflect.TypeOf(MemoryMqInput{})
	Type["socket"] = reflect.TypeOf(SocketInput{})
}
