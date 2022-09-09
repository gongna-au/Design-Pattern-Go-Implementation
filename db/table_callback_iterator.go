package db

import (
	"fmt"
)

type callbacktableIteratorImpl struct {
	rs []record
}

type Callback func(*record)

func PrintRecord(record *record) {
	fmt.Printf("%+v\n", record)
}

func (c *callbacktableIteratorImpl) Iterator(callback Callback) {
	go func() {
		for _, re := range c.rs {
			callback(&re)
		}
	}()

}

func (r *callbacktableIteratorImpl) HasNext() bool {
	return true
}

// 关键点: 在Next函数中取出下一个记录，并转换成客户端期望的对象类型，记得增加cursor
func (r *callbacktableIteratorImpl) Next(next interface{}) error {

	return nil
}

//用工厂模式来创建我们这个复杂的迭代器
type callbackTableIteratorFactory struct {
}

func NewcallbackTableIteratorFactory() *complexTableIteratorFactory {
	return &complexTableIteratorFactory{}

}

func (c *callbackTableIteratorFactory) Create(table *Table) TableIterator {
	var res []record
	for _, r := range table.records {
		res = append(res, r)
	}
	return &complextableIteratorImpl{
		rs: res,
	}
}
