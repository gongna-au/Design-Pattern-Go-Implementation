package flyweight

// Flyweight 模式只是一种优化。
// 在应用它之前，
// 1.确保程序确实存在与内存中同时存在大量相似对象相关的 RAM 消耗问题。
// 2.Flyweight 类包含(可以在多个对象之间共享的)原始对象状态的一部分。
// 3.享元类要作为context的属性或者方法的参数。同一个享元对象可以在许多不同的上下文中使用。
// 4.存储在享元中的状态称为内在状态。享元类需要一个属性来存储——内在状态
// 5.传递给享元方法的状态称为外部状态享元类需要一个方法来接收外部状态
// 6.Context 需要一个属性来存储外部状态
// 7.Context 需要一个属性来存储享元对象，当上下文和一个享元对象配对的时，代表原始对象最完整的状态u

type Flyweight struct {
	// repeatingState 内部状态
	repeatingState int
}

// uniqueState 外部状态
func (f *Flyweight) Operation(uniqueState int) {
	f.repeatingState = uniqueState
}

type FlyweightFactory struct {
	cache []Flyweight
}

func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{}
}
func (f *FlyweightFactory) GetFlyweight(repeatingState int) Flyweight {
	return f.cache[repeatingState]
}

type Context struct {
	// 存储享元的外部状态
	uniqueState int
	// 存储享元对象
	Flyweight Flyweight
}

func NewContext(uniqueState int) *Context {
	c := &Context{
		uniqueState: uniqueState,
	}
	c.Flyweight = NewFlyweightFactory().GetFlyweight(uniqueState)
	return c
}
