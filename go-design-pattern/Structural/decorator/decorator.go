package decorator

type IComponent interface {
	execute()
}

type ConcreteComponent struct {
}

func (c *ConcreteComponent) execute() {

}

// 供给参数为IComponent类型的地方调用
func NewConcreteComponent() *ConcreteComponent {
	return &ConcreteComponent{}
}

// 抽象的装饰器
type BaseDecorator struct {
	component IComponent
}

// 供给子类调用使用
func NewBaseDecorator(component IComponent) *BaseDecorator {
	return &BaseDecorator{
		component: component,
	}
}

func (c *BaseDecorator) execute() {
	c.component.execute()
}

// 具体的装饰器材继承BaseDecorator
type ConcreteDecorator struct {
	*BaseDecorator
}

func NewConcreteDecorator(component IComponent) *ConcreteDecorator {
	return &ConcreteDecorator{
		BaseDecorator: NewBaseDecorator(component),
	}
}

func (c *ConcreteDecorator) execute() {
	// do something
	c.component.execute()
}

func Client() {
	d := NewConcreteDecorator(NewConcreteComponent())
	d.execute()
}
