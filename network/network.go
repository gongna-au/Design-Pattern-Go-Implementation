package network

import (
	"errors"
	"sync"
)

/*
单例模式
*/
// 往网络的作用就是在多个地址上发起多个socket监听
// 所以我们需要一个map结构来存储这种状态
type network struct {
	sockets sync.Map
}

// 懒汉版单例模式
var networkInstance = &network{
	sockets: sync.Map{},
}

func GetnetworkInstance() *network {
	return networkInstance
}

// Listen 在endpoint指向地址上起监听 endpoint资源是暴露一个服务的ip地址和port的列表。
// 监听的本质就是把目的地址和对目的地址的连接添加到网络的map存储结构当中
// 用于socktImpl来调用
// 这里的endpoint 是网络包里的目的地址，而socket里面存储的是目的地址对应的socket
func (n *network) Listen(endpoint Endpoint, socket Socket) error {
	if _, ok := n.sockets.Load(endpoint); ok {
		return errors.New("ErrEndpointAlreadyListened")
	}
	n.sockets.Store(endpoint, socket)
	return nil
}

// 用于socktImpl来调用
func (n *network) Disconnect(endpoint Endpoint) {
	n.sockets.Delete(endpoint)

}

// 用于socktImpl来调用
func (n *network) DisconnectAll() {
	n.sockets = sync.Map{}

}

// 网络的发送作用就是 向目的地址发送包裹
// 包裹中含有目的地址和数据
// 应该先在map中根据目的地址获取到连接，然后才能向连接发送数据
// 向连接发送数据的本质就是 这个连接去接收到数据
func (n *network) Send(packet *Packet) error {
	con, okc := n.sockets.Load(packet.Dest())
	socket, oks := con.(Socket)
	if !okc || !oks {
		return errors.New("ErrConnectionRefuse")
	}
	go socket.Receive(packet)
	return nil
}

/*
// 其余单例模式实现

type network struct {}
var once sync.Once
var netnetworkInstance *network
func GetnetworkInstance() *network {
	once.Do(func (){
		netnetworkInstance=&network {

		}

	})
	return netnetworkInstance

}

*/
