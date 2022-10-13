### 装饰器模式的结构

#### 1.接口 IComponent 声明了包装器和被包装对象的通用接口。

```go

type IComponent interface {
	execute()
}

```

#### 2.具体的组件实现了 IComponent 接口

```go
type ConcreteComponent struct {
}

func (c *ConcreteComponent) execute() {

}

func NewConcreteComponent() *ConcreteComponent {
	return &ConcreteComponent{}
}

```

#### 3.抽象的 BaseDecorator 装饰器实现了 IComponent 接口，BaseDecorator 组合了 IComponent 接口

```go
type BaseDecorator struct {
	component IComponent
}

func NewBaseDecorator(component IComponent) *BaseDecorator {
	return &BaseDecorator{
		component: component,
	}
}

func (c *BaseDecorator) execute() {
	c.component.execute()
}

```

#### 4.具体的的装饰器继承抽象的继承器（具体的装饰器就可以横向扩展抽象的装饰器的功能）

```go

type ConcreteDecorator struct {
	BaseDecorator
}

func NewConcreteDecorator(component IComponent) *BaseDecorator {
	return &BaseDecorator{
		component: component,
	}
}

func (c *ConcreteDecorator) execute() {
	// do something
	c.component.execute()
}

```

#### 5.客户端调用具体的装饰器进行装饰

```go

func Client() {
	d := NewConcreteDecorator(NewConcreteComponent())
	d.execute()
}

```
