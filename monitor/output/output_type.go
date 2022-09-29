package output

import (
	"github.com/Design-Pattern-Go-Implementation/monitor/plugin"
	"reflect"
)

// Type output插件类型变量
var Type = make(plugin.Types)

func init() {
	Type["memory_db"] = reflect.TypeOf(MemoryDbOutput{})
}
