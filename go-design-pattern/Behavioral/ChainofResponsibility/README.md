### 责任链行为设计模式的结构：

> 解耦请求的发送和接收，让多个接收对象都有机会处理这个请求。请求在多个串联的对象中传递，直到有对象能处理它。

#### 1.责任链接口一定实现了两个函数：SetNext(h IHandler)将“本“接口作为函数参数，Hadle(request string)接收外界传入的数据

```go
type IHandler interface {
	SetNext(h IHandler)
	Hadle(request string)
}

```

#### 2.BaseHandler 组合 IHandler 接口，目的是存储传入的同样实现了 IHandler 接口的实例，并在自己的处理逻辑执行完毕的时候，调用保存的这个实例的方法

```go
type BaseHandler struct {
	Next IHandler
}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

```

#### 3.BaseHandler 实现 IHandler 接口，其中重要的一个方法就是设置调用链上的下一个实例，另外一个方法要实现的功能就是：接收外界传入的数据，处理数据，如果下一个实例不为空，那么就调用下一个实例的"该"方法。使得链条传递下去，而不是执行到第一个就结束了。

```go

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
```

#### 4.客户端提供链上的所有实例，要做的就是调用已经有的函数和方法，组装链，调用链上的第一个实例，就可以完整的调用整个链

```go
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

```
