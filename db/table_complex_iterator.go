package db

// 关键点: 定义迭代器的另外一种实现的实现
// tableIteratorImpl 迭代器接口公共实现类 用来实现遍历表
type complextableIteratorImpl struct {
	rs []record
}

// 关键点1: 定义迭代器创建方法，返回只能接收的channel类型
func (c *complextableIteratorImpl) Iterator() <-chan record {
	// 关键点2: 创建一个无缓冲的channel
	ch := make(chan record)
	// 关键点3: 另起一个goroutine往channel写入记录，如果接收端还没开始接收，会阻塞住
	go func() {
		for _, re := range c.rs {
			ch <- re
		}
		close(ch)
	}()
	return ch
}

// 关键点5: 在HasNext函数中的判断是否已经遍历完所有记录
func (r *complextableIteratorImpl) HasNext() bool {
	return true
}

// 关键点: 在Next函数中取出下一个记录，并转换成客户端期望的对象类型，记得增加cursor
func (r *complextableIteratorImpl) Next(next interface{}) error {

	return nil
}

//用工厂模式来创建我们这个复杂的迭代器
type complexTableIteratorFactory struct {
}

func NewcomplexTableIteratorFactory() *complexTableIteratorFactory {
	return &complexTableIteratorFactory{}

}
func (c *complexTableIteratorFactory) Create(table *Table) TableIterator {
	var res []record
	for _, r := range table.records {
		res = append(res, r)
	}

	return &complextableIteratorImpl{
		rs: res,
	}

}
