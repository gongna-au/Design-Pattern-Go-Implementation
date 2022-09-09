package network

//一个网络包裹包括一个源ip地址和端口 和目的地址ip和端口
type Packet struct {
	src     Endpoint
	dest    Endpoint
	payload interface{}
}

func NewPacket(src, dest Endpoint, payload interface{}) *Packet {
	return &Packet{
		src:     src,
		dest:    dest,
		payload: payload,
	}
}

//返回源地址
func (p Packet) Src() Endpoint {
	return p.src
}

//返回目的地址
func (p Packet) Dest() Endpoint {
	return p.dest
}

func (p Packet) Payload() interface{} {
	return p.payload
}
