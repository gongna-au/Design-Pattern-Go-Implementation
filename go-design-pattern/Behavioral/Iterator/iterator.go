package iterator

// 迭代器模式的核心：是将集合的遍历行为提取到称为迭代器的单独对象中。
// 向集合中添加越来越多的遍历算法逐渐模糊了它的主要职责，即高效的数据存储。此外，某些算法可能会针对特定应用程序进行定制，因此将它们包含在通用集合类中会很奇怪。
// 另一方面，应该与各种集合一起使用的客户端代码可能甚至不关心它们如何存储元素。但是，由于集合都提供了访问其元素的不同方式，因此您别无选择，只能将代码耦合到特定的集合类。
// 需要一种特殊的方式来遍历集合，只需创建一个新的迭代器类，而无需更改集合或客户端。
// IIteratorCollection 迭代器需要实现的接口
type IIterator interface {
	//NewIterator(IIteratorCollection)IIterator
}

type ConcreteIterator struct {
}

func NewConcreteIterator() *ConcreteIterator {
	return &ConcreteIterator{}
}

// IIteratorCollection可以被遍历的集合的接口
// 集合实现的接口
// 集合在自己需要被遍历的时候调用自己的CreateIterator方法得到一个指针，然后调用指针的方法来得到元素
// ConcreteIterator组合『集合接口』并通过具体实现了『集合接口的具体的类』创建出来
// 对于具体的集合类而言，返回的是IIterator
type IIteratorCollection interface {
	CreateIterator() IIterator
}

type ConcreteIteratorCollection struct {
}

func (c *ConcreteIteratorCollection) CreateIterator() IIterator {

}
