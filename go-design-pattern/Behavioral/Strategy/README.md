### 策略模式的结构：

#### 1.每一种的算法都有属于自己的类

```go
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

```

#### 2. 抽象出这些算法类通用的接口

```go
type IStrategy interface {
	RoutePlanning(start string, end string)
}
```

#### 3.业务上下文 Context 需要做的就是组合这个接口（存在一个类型为 IStrategy 的属性），在需要具体的算法的时候，把具体算法对应的类的对象传进去，

```go

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
```

#### 4.业务上下文 Context 在需要的调用具体的 Strategy 的方法

```go
// Context 一定有个方法来调用具体的Strategy 的方法
func (c *Context) Execute(start string, end string) {
	c.Strategy.RoutePlanning(start, end)
}
```

#### 5.客户端

```go

func Client() {
	start := "华中师范大学"
	end := "武汉大学"
	b := NewBike()
	context := NewContext()
	context.SetStrategy(b)
	context.Execute(start, end)
}
```
