package chainofresponsibility

import "fmt"

// gin 的中间件实现原理
// 责任链模式在实际使用时
// 更多的是另一种变体，请求会被串联的多个对象依次处理，而不是被某个对象处理后就停止处理了。类似这样
// 但是下面的这种耦合程度比较高，意味着你要把所有的逻辑加到func(*Request) bool函数
// 意味着这个函数可能很长
type Request struct {
	url      string
	handlers []HandlerFunc
}

func NewRequest(url string) *Request {
	return &Request{}
}

type HandlerFunc func(*Request) bool

//通过 Request.Use() 方法进行添加HandlerFunc
func (r *Request) Use(h HandlerFunc) {
	r.handlers = append(r.handlers, h)
}

//通过 Request.Run() 方法进行调用
func (r *Request) Run() {
	for _, v := range r.handlers {
		if !v(r) {
			break
		}
	}
}

func GinClient() {

	r := NewRequest("http://www.baidu.com/index.html?name=mo&age=25#dowell")
	r.Use(func(r *Request) bool {
		fmt.Println("first middware function")
		return true
	})

	r.Use(func(r *Request) bool {
		fmt.Println("second middware function")
		return true
	})

	r.Use(func(r *Request) bool {
		fmt.Println("third middware function")
		return true
	})

	r.Use(func(r *Request) bool {
		fmt.Println("four middware function")
		return true
	})
	r.Run()

}
