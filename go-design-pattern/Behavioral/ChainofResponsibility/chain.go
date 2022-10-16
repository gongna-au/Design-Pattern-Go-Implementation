package chainofresponsibility

type IHandler interface {
	SetNext(h IHandler)
	Hadle(request string)
}

type BaseHandler struct {
	Next IHandler
}

func (b *BaseHandler) SetNext(h IHandler) {
	b.Next = h
}

func (b *BaseHandler) Hadle(request string) {
	// BaseHandler process the request
	// BaseHandler may or may not change the request
	request = request + "baseHandler1"
	if b.Next != nil {
		b.Next.Hadle(request)
	}
}

type ConcreteHandler struct {
	*BaseHandler
	request string
}

func NewConcreteHandler() *ConcreteHandler {
	return &ConcreteHandler{}
}

// SetRequest 保存Request，然后交给扩展的功能进行处理
func (c *ConcreteHandler) SetRequest(request string) {
	c.request = request
}

// RequestAddName Extend function and override parent class function 扩展功能
func (c *ConcreteHandler) RequestAddName() {
	c.request = c.request + "name"

}

// RequestAddColor Extend function and override parent class function 扩展功能
func (c *ConcreteHandler) RequestAddColor() {
	c.request = c.request + "color"
}

// RequestAddAge Extend function and override parent class function 扩展功能
func (c *ConcreteHandler) RequestAddAge() {
	c.request = c.request + "age"
}

func (c *ConcreteHandler) GetRequest() string {
	return c.request
}

// 重写父类函数
func (c *ConcreteHandler) Hadle(request string) {
	c.SetRequest(request)
	c.RequestAddName()
	c.RequestAddAge()
	c.RequestAddColor()
	c.SetNext(NewConcreteHandler())
	c.Next.Hadle(c.GetRequest())
}
