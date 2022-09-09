package http

import (
	"errors"
	"github.com/Design-Pattern-Go-Implementation/network"
)

// Handler HTTP请求处理接口
type Handler func(req *Request) *Response

// Server Http服务器
type Server struct {
	socket        network.Socket
	localEndpoint network.Endpoint
	routers       map[Method]map[Uri]Handler
}

func NewServer(socket network.Socket) *Server {
	server := &Server{
		socket:  socket,
		routers: make(map[Method]map[Uri]Handler),
	}
	server.socket.AddListener(server)
	return server
}

// 实现 Handle 方法才能被添加到listeners中
// Server处理的是请求数据包
// Client处理的是响应数据包
// 请求数据包的路径   Client 发出请求数据包  ——>  Network拿到请求数据包  ———> Network把请求数据包给到 Server （Server拿到请求数据包）
// 响应数据包的处理   Sever拿到请求处理包处理得到响应数据包  ————>  Server 把响应数据包给到Network ————> network 拿到响应数据包 然后把响应数据包给Client ————> client拿到响应数据包
func (s *Server) Handle(packet *network.Packet) error {
	req, ok := packet.Payload().(*Request)
	if !ok {
		return errors.New("invalid packet, not http request")
	}
	if req.IsInValid() {
		resp := ResponseOfId(req.ReqId()).
			AddStatusCode(StatusBadRequest).
			AddProblemDetails("uri or method is invalid")
		return s.socket.Send(network.NewPacket(packet.Dest(), packet.Src(), resp))
	}

	router, ok := s.routers[req.Method()]
	if !ok {
		resp := ResponseOfId(req.ReqId()).
			AddStatusCode(StatusMethodNotAllow).
			AddProblemDetails(StatusMethodNotAllow.Details)
		return s.socket.Send(network.NewPacket(packet.Dest(), packet.Src(), resp))
	}

	var handler Handler
	//得到所有的路由，然后把所有的路由和请求网络包中的携带的要请求的路由进行匹配
	for u, h := range router {
		if req.Uri().Contains(u) {
			handler = h
			break
		}
	}

	if handler == nil {
		resp := ResponseOfId(req.ReqId()).
			AddStatusCode(StatusNotFound).
			AddProblemDetails("can not find handler of uri")
		return s.socket.Send(network.NewPacket(packet.Dest(), packet.Src(), resp))
	}

	resp := handler(req)
	return s.socket.Send(network.NewPacket(packet.Dest(), packet.Src(), resp))
}

// 客户端的监听和服务端的监听的区别在于：
// 客户端监听是把自己的socketImpl 和 自己的ip地址对应的endpoint 添加到网络中，当客户端发送网络数据包的时候， network 根据网络数据包里面的 destEndpoint
// 查询map 得到socketImpl   network 把网络数据包交给服务端口的socketImpl
// 服务端的监听 就是根据ip 和端口 port把server的localEndpoint给赋值
// 服务器和客户端不同的是服务器是通过Start() 才把本地的localEndpoint 添加到网络中，
// 先存储再添加的方式可以使得服务器材自由控制是否开启
func (s *Server) Listen(ip string, port int) *Server {
	s.localEndpoint = network.EndpointOf(ip, port)
	return s
}

func (s *Server) Start() error {
	return s.socket.Listen(s.localEndpoint)
}

func (s *Server) Shutdown() {
	s.socket.Close(s.localEndpoint)
}

// server 下的network.Endpoint （被观察者得到网络数据包肯定会把网络数据包层层往上抛）
// 抛给 Server对应的socketImpl 肯定会通知 它的listeners （server）来处理，因为会得到关于 request的数据，那么当然就要解析数据
// 服务器的路由肯定是需要你去自己设置然后初始化的
// Get （） Post（） Put（）  Delete（） 需要服务端人为的设置路由和路由对应的处理函数  ，处理函数 一定是以req *Request 为参数，且返回一个Response
func (s *Server) Get(uri Uri, handler Handler) *Server {
	if _, ok := s.routers[GET]; !ok {
		s.routers[GET] = make(map[Uri]Handler)
	}
	s.routers[GET][uri] = handler
	return s
}

func (s *Server) Post(uri Uri, handler Handler) *Server {
	if _, ok := s.routers[POST]; !ok {
		s.routers[POST] = make(map[Uri]Handler)
	}
	s.routers[POST][uri] = handler
	return s
}

func (s *Server) Put(uri Uri, handler Handler) *Server {
	if _, ok := s.routers[PUT]; !ok {
		s.routers[PUT] = make(map[Uri]Handler)
	}
	s.routers[PUT][uri] = handler
	return s
}

func (s *Server) Delete(uri Uri, handler Handler) *Server {
	if _, ok := s.routers[DELETE]; !ok {
		s.routers[DELETE] = make(map[Uri]Handler)
	}
	s.routers[DELETE][uri] = handler
	return s
}
