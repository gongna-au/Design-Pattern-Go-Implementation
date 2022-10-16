package strategy

import "fmt"

// 核心：提取一系列不同的算法
// 把每个算法都提取到单独的类当中，并使得他们的对象之间可以进行替换
// 实现：肯定要定义接口来实现，并且这些这些类实现了相同的方法
// 业务函数需要做的就是切换他们

type IStrategy interface {
	RoutePlanning(start string, end string)
}

// 关于汽车的路线规划
type Bus struct {
}

func NewBus() *Bus {
	return &Bus{}
}

func (b *Bus) RoutePlanning(start string, end string) {
	fmt.Println("Bus route is planned")
}

// 关于小汽车的路线规划
type Car struct {
}

func NewCar() *Car {
	return &Car{}
}

func (c *Car) RoutePlanning(start string, end string) {
	fmt.Println("Car route is planned")
}

// 关于自行车的路线规划
type Bike struct {
}

func NewBike() *Bike {
	return &Bike{}
}

func (b *Bike) RoutePlanning(start string, end string) {
	fmt.Println("Bike route is planned")
}

type Context struct {
	Strategy IStrategy
}

func NewContext() *Context {
	return &Context{}
}

// Context  一定有个方法是去设置（改变）具体的Strategy
func (c *Context) SetStrategy(i IStrategy) {
	c.Strategy = i
}

// Context 一定有个方法来调用具体的Strategy 的方法
func (c *Context) Execute(start string, end string) {
	c.Strategy.RoutePlanning(start, end)
}

func Client() {
	start := "华中师范大学"
	end := "武汉大学"
	b := NewBike()
	context := NewContext()
	context.SetStrategy(b)
	context.Execute(start, end)
}
