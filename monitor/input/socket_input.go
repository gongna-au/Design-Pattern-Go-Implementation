package input

import (
	"github.com/Design-Pattern-Go-Implementation/monitor/plugin"
	"github.com/Design-Pattern-Go-Implementation/network"
	"sync/atomic"
)

// SocketInput 具体的插件实现了Plugin 接口
type SocketInput struct {
	socket      network.Socket
	endpoint    network.Endpoint
	packets     chan *network.Packet
	isUninstall uint32
}

func (s *SocketInput) Install() {
	s.socket = network.DefaultSocket()
	s.packets = make(chan *network.Packet, 10000)
	s.socket.AddListener(s)
	s.socket.Listen(s.endpoint)
}

func (s *SocketInput) Uninstall() {
	close(s.packets)
	atomic.StoreUint32(&s.isUninstall, 1)
	s.socket.Close(s.endpoint)
}

func (s *SocketInput) SetContext(ctx plugin.Context) {
	ip, ok := ctx.GetString("ip")
	if !ok {
		return
	}
	port, ok := ctx.GetInt("port")
	if !ok {
		return
	}
	s.endpoint = network.EndpointOf(ip, port)
}

// 具体的Input动作
func (s *SocketInput) Input() (*plugin.Event, error) {
	packet, ok := <-s.packets
	if !ok {
		return nil, plugin.ErrPluginUninstalled
	}
	event := plugin.NewEvent(packet.Payload())
	event.AddHeader("peer", packet.Src().String())
	return event, nil
}

//
func (s *SocketInput) Handle(packet *network.Packet) error {
	if s.socket == nil {
		return plugin.ErrPluginNotInstalled
	}
	if atomic.LoadUint32(&s.isUninstall) == 1 {
		return plugin.ErrPluginUninstalled
	}
	s.packets <- packet
	return nil
}
