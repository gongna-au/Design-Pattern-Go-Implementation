package chainofresponsibility

import (
	"fmt"
	"testing"
)

func TestChain(t *testing.T) {
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
