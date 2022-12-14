package sidecar

import (
	"github.com/Design-Pattern-Go-Implementation/mq"
	"github.com/Design-Pattern-Go-Implementation/network"
)

// AllInOneFactory 具备所有功能的sidecar工厂
type AllInOneFactory struct {
	producer mq.Producible
}

func NewAllInOneFactory(producer mq.Producible) *AllInOneFactory {
	return &AllInOneFactory{producer: producer}
}

func (a AllInOneFactory) Create() network.Socket {
	return NewAccessLogSidecar(NewFlowCtrlSidecar(network.DefaultSocket()), a.producer)
}
