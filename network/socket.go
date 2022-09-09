package network

/*
观察者模式
*/
// socketListener 需要作出反应，就是向上获取数据package network
// SocketListener Socket报文监听者
// SocketListener 应该是一个客户端Client
type SocketListener interface {
	Handle(packet *Packet) error
}

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

// 被观察者在未知的情况下应该先定义一个接口来代表观察者们
// 被观察者往往应该持有观察者列表
// socketImpl Socket的默认实现
type socketImpl struct {
	// 关键点4: 在Subject中持有Observer的集合
	listeners []SocketListener
}

// Listen 在endpoint指向地址上起监听 endpoint资源是暴露一个服务的ip地址和port的列表。（endpoint来自与包裹里面的目的地址）
func (s *socketImpl) Listen(endpoint Endpoint) error {
	return GetnetworkInstance().Listen(endpoint, s)
}

func (s *socketImpl) Close(endpoint Endpoint) {
	GetnetworkInstance().Disconnect(endpoint)

}

func (s *socketImpl) Send(packet *Packet) error {
	return GetnetworkInstance().Send(packet)
}

// 关键点: 为Subject定义注册Observer的方法(为被观察者提供添加观察者的方法)
func (s *socketImpl) AddListener(listener SocketListener) {
	s.listeners = append(s.listeners, listener)

}

// 关键点: 当Subject状态变更时，遍历Observers集合，调用它们的更新处理方法
// 当被观察者的状态发生变化的时候，需要遍历观察者的列表来调用观察者的行为
// 被观察者一定有一个函数，用来在自己的状态改变时通知观察者们进行一系列的行为
// 这里的状态改变（就是被观察者收到外界来的实参）
func (s *socketImpl) Receive(packet *Packet) {
	for _, listener := range s.listeners {
		listener.Handle(packet)
	}

}
