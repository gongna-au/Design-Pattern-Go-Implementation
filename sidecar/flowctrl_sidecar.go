package sidecar

import (
	"github.com/Design-Pattern-Go-Implementation/http"
	"github.com/Design-Pattern-Go-Implementation/network"
	"github.com/Design-Pattern-Go-Implementation/sidecar/flowctrl"
)

// 关键点1: 定义被装饰的抽象接口
// Socket 网络通信Socket接口

/*
type Socket interface {
     // Listen 在endpoint指向地址上起监听
     Listen(endpoint Endpoint) error
     // Close 关闭监听
     Close(endpoint Endpoint)
     // Send 发送网络报文
     Send(packet *Packet) error
     // Receive 接收网络报文
     Receive(packet *Packet)
     // AddListener 增加网络报文监听者
     AddListener(listener SocketListener)
}

*/
// 关键点2: 提供一个默认的基础实现
/*
type socketImpl struct {
     listener SocketListener
}
*/
// 关键点: 定义被装饰的抽象接口(需要装饰什么就先提取出来一个抽象的接口，为什么必须是接口而不是实例？)
// 因为我们写的sidecar可能是针对一组业务逻辑，而这组业务逻辑往往提供很多不同的实现
type FlowCtrlSidecar struct {
	// 关键点4: 装饰器持有被装饰的抽象接口作为成员属性
	socket network.Socket
	ctx    *flowctrl.Context
}

// 关键点7: 定义装饰器的工厂方法，入参为被装饰接口
// 调用方式:NewFlowCtrlSidecar(network.DefaultSocket())
func NewFlowCtrlSidecar(socket network.Socket) *FlowCtrlSidecar {
	return &FlowCtrlSidecar{
		socket: socket,
		ctx:    flowctrl.NewContext(),
	}
}

func (f *FlowCtrlSidecar) Listen(endpoint network.Endpoint) error {
	return network.GetnetworkInstance().Listen(endpoint, f)
}

// 关键点6: 不需要扩展功能的方法，直接调用被装饰接口的原生方法即可
// 对接口的装饰器为什么需要持有被装饰接口的属性，这个是为了对接口下面的方法进行扩展
// 装饰器写一些函数，这些函数的类型和被装饰接口的类型一模一样，不同与子类继承父类重写父类方法
// 而是写一个函数（把被装饰接口下的函数包起来），在这个函数的前部分加上扩展了功能代码
//（把被装饰接口下的函数包起来）就是装饰器调用自己保存的(被装饰器的成员变量调用对应的函数)
// 不需要扩展的功能就是直接包起来，函数前部分不增加功能代码
func (f *FlowCtrlSidecar) Close(endpoint network.Endpoint) {
	f.socket.Close(endpoint)
}

func (f *FlowCtrlSidecar) Send(packet *network.Packet) error {
	return f.socket.Send(packet)
}

// 关键点5: 对于需要扩展功能的方法，新增扩展功能
// 新增的拓展功能如何实现？
// 拓展的功能就是相当与给持有的被装饰的抽象接口增加方法
func (f *FlowCtrlSidecar) Receive(packet *network.Packet) {
	httpReq, ok := packet.Payload().(*http.Request)
	// 如果不是HTTP请求，则不做流控处理
	if !ok {
		f.socket.Receive(packet)
		return
	}
	// 流控后返回429 Too Many Request响应
	// f.ctx.TryAccept() 是把传入的f.ctx.reqCount 属性加1
	// 然后尝试f.ctx.curState 状态改版
	if !f.ctx.TryAccept() {
		httpResp := http.ResponseOfId(httpReq.ReqId()).
			AddStatusCode(http.StatusTooManyRequest).
			AddProblemDetails("enter flow ctrl state")
		f.socket.Send(network.NewPacket(packet.Dest(), packet.Src(), httpResp))
		return
	}
	f.socket.Receive(packet)
}

func (f *FlowCtrlSidecar) AddListener(listener network.SocketListener) {
	f.socket.AddListener(listener)
}

// 关键点8: 使用时，通过装饰器的工厂方法，把所有装饰器和被装饰者串联起来
// 装饰器的工厂方法 就是把FlowCtrlSidecar 一个接一个连接起来

/* func (a AllInOneFactory) Create() network.Socket {
	return NewAccessLogSidecar(NewFlowCtrlSidecar(network.DefaultSocket()), a.producer)
} */
