package sidecar

import "github.com/Design-Pattern-Go-Implementation/network"

/*
工厂方法模式
*/

// Factory Sidecar工厂接口
type Factory interface {
	Create() network.Socket
}
