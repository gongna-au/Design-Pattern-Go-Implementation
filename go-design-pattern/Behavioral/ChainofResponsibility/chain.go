package chainofresponsibility

import "fmt"

type IHandler interface {
	SetNext(h IHandler)
	Hadle(request string)
}

type BaseHandler struct {
	Next IHandler
}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
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

type ConcreteHandler1 struct {
	*BaseHandler
	request string
}

func NewConcreteHandler1() *ConcreteHandler1 {
	return &ConcreteHandler1{
		BaseHandler: NewBaseHandler(),
	}
}

// SetRequest 保存Request，然后交给扩展的功能进行处理
func (c *ConcreteHandler1) SetRequest(request string) {
	c.request = request
}

// RequestAddName Extend function and override parent class function 扩展功能
func (c *ConcreteHandler1) RequestAddName() {
	c.request = c.request + " name"

}

// RequestAddColor Extend function and override parent class function 扩展功能
func (c *ConcreteHandler1) RequestAddColor() {
	c.request = c.request + " color"
}

// RequestAddAge Extend function and override parent class function 扩展功能
func (c *ConcreteHandler1) RequestAddAge() {
	c.request = c.request + " age"
}

func (c *ConcreteHandler1) GetRequest() string {
	return c.request
}

// 重写父类函数
func (c *ConcreteHandler1) Hadle(request string) {
	c.SetRequest(request)
	c.RequestAddName()
	c.RequestAddColor()
	c.RequestAddAge()
	if c.Next != nil {
		c.Next.Hadle(c.GetRequest())
	}

}

type ConcreteHandler2 struct {
	*BaseHandler
	request string
}

func NewConcreteHandler2() *ConcreteHandler2 {
	return &ConcreteHandler2{
		BaseHandler: NewBaseHandler(),
	}
}

// SetRequest 保存Request，然后交给扩展的功能进行处理
func (c *ConcreteHandler2) SetRequest(request string) {
	c.request = request
}

// RequestAddTrace Extend function and override parent class function 扩展功能
func (c *ConcreteHandler2) RequestAddTrace() {
	c.request = c.request + " trace"

}

// RequestAddMetric Extend function and override parent class function 扩展功能
func (c *ConcreteHandler2) RequestAddMetric() {
	c.request = c.request + " metric"
}

func (c *ConcreteHandler2) GetRequest() string {
	return c.request
}

// 重写父类函数
func (c *ConcreteHandler2) Hadle(request string) {
	c.SetRequest(request)
	c.RequestAddTrace()
	c.RequestAddMetric()
	if c.Next != nil {
		c.Next.Hadle(c.GetRequest())
	}

}

func Client() {
	h1 := NewConcreteHandler1()
	fmt.Println("h1 create")
	h2 := NewConcreteHandler2()
	fmt.Println("h2 create")
	h3 := NewConcreteHandler1()
	fmt.Println("h3 create")
	h4 := NewConcreteHandler2()
	fmt.Println("h4 create")
	h1.SetNext(h2)
	fmt.Println("h1 set next h2")
	h2.SetNext(h3)
	fmt.Println("h2 set next h3")
	h3.SetNext(h4)
	fmt.Println("h3 set next h4")
	h1.Hadle("request")
	fmt.Println(h4.GetRequest())
}
