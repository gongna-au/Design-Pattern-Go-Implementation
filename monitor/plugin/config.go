package plugin

import (
	"reflect"
	"strconv"
)

// 待补充
type Types map[string]reflect.Type

// Config 插件配置抽象接口
type Config interface {
	Load(conf string) error
}

// Context 插件配置上下文
type Context map[string]string

func EmptyContext() Context {
	return make(map[string]string)
}

func (c Context) Add(key string, value string) {
	c[key] = value
}

func (c Context) GetString(key string) string {
	v, ok := c[key]
	if ok {
		return v

	} else {
		return ""
	}

}

func (c Context) GetInt(key string) (int, bool) {
	val, ok := c[key]
	if !ok {
		return 0, false
	}
	if iVal, err := strconv.Atoi(val); err == nil {
		return iVal, true
	}
	return 0, false
}
