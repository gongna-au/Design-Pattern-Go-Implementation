package http

import (
	"errors"
	"github.com/Design-Pattern-Go-Implementation/network"
	"math/rand"
	"time"
)

// 观察者包含被观察者就可以封装被观察者，调用被观察者
// 一般来说，观察者往往是用户，所以如果观察者存储有被观察者，那么就可以调用被观察者的接口实现一系列操作，
// 而对对于网络来说，更像是两个被观察者在面对面交谈，而实际用户（观察者）因为保存有被观察者因而看起来像很多观察者面对面交流
type Client struct {
	// 接收网络数据包并且
	socket network.Socket
	// 把处理的结果写入到一个channel中，因为处理结果是有数据的
	respChan chan *Response
	// 代表者自己的的ip地址和端口
	localEndpoint network.Endpoint
}

// 通过本机的ip 以及随即生成一个端口，代表本机的这个端口下的程序
func NewClient(socket network.Socket, ip string) (*Client, error) {
	// 一个观察者肯定有一个被观察者需要他去观察
	// 一个client 肯定有一个 ip 代表自己要访问的
	// 随机端口，从10000 ～ 19999
	endpoint := network.EndpointOf(ip, int(rand.Uint32()%10000+10000))
	client := &Client{
		socket:        socket,
		localEndpoint: endpoint,
		respChan:      make(chan *Response),
	}
	// 一个观察者开始观察一个（被观察者）的时候，
	// 也就意味着被观察者的监听列表肯定要把这个观察者加入它的列表
	// 二者是同步的
	client.socket.AddListener(client)
	// 把本机器的socketImpl 添加到全局唯一一个的且被共享的网络实例
	if err := client.socket.Listen(endpoint); err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) Close() {
	//从全局的网络中删除
	c.socket.Close(c.localEndpoint)
	close(c.respChan)
}

// 底层调用network的Send 然后网络是根据网络包中目的地址 一下子得到目的地址对应的
func (c *Client) Send(dest network.Endpoint, req *Request) (*Response, error) {
	// 制作网络包 网络包包含着目的endpoint  通过目的endpoint可以在网络中查到对应的socketImpl（被观察者）
	// req是携带的数据
	packet := network.NewPacket(c.localEndpoint, dest, req)
	// 通过底层调用network.Send()
	// network.Send()就是根据网络数据包的目的地址得到对应的socketImpl
	// 然后把数据发给socketImpl  ，socketImpl 一旦接收到数据，就是调用自己listeners的也就是client去处理
	err := c.socket.Send(packet)
	if err != nil {
		return nil, err
	}
	// 发送请求后同步阻塞等待响应
	select {
	case resp, ok := <-c.respChan:
		if ok {
			return resp, nil
		}
		errResp := ResponseOfId(req.ReqId()).AddStatusCode(StatusInternalServerError).
			AddProblemDetails("connection is break")
		return errResp, nil
	case <-time.After(time.Second * time.Duration(3)):
		// 超时时间为3s
		resp := ResponseOfId(req.ReqId()).AddStatusCode(StatusGatewayTimeout).
			AddProblemDetails("http server response timeout")
		return resp, nil
	}
}

//
func (c *Client) Handle(packet *network.Packet) error {
	resp, ok := packet.Payload().(*Response)
	if !ok {
		return errors.New("invalid packet, not http response")
	}
	c.respChan <- resp
	return nil
}
