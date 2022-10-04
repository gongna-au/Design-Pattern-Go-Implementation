package creation

import (
	"fmt"
	"reflect"
)

// DI容器
// Container DI 容器
type Container struct {
	// 假设一种类型只能有一个 provider 提供
	providers map[reflect.Type]provider

	// 缓存以生成的对象
	results map[reflect.Type]reflect.Value
}
type provider struct {
	value reflect.Value

	params []reflect.Type
}

// New 创建一个容器
func New() *Container {
	return &Container{
		providers: map[reflect.Type]provider{},
		results:   map[reflect.Type]reflect.Value{},
	}
}
