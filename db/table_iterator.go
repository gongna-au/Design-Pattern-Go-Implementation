package db

import (
	"math/rand"
	"sort"
	"time"
)

type TableIterator interface {
	HasNext() bool
	Next(next interface{}) error
}

// 关键点: 定义迭代器接口的实现
// tableIteratorImpl 迭代器接口公共实现类 用来实现遍历表
type tableIteratorImpl struct {
	// 关键点3: 定义一个集合存储待遍历的记录，这里的记录已经排序好或者随机打散
	records []record
	// 关键点4: 定义一个cursor游标记录当前遍历的位置
	cursor int
}

// 关键点5: 在HasNext函数中的判断是否已经遍历完所有记录
func (r *tableIteratorImpl) HasNext() bool {
	return r.cursor < len(r.records)
}

// 关键点: 在Next函数中取出下一个记录，并转换成客户端期望的对象类型，记得增加cursor
func (r *tableIteratorImpl) Next(next interface{}) error {
	record := r.records[r.cursor]
	r.cursor++
	if err := record.convertByValue(next); err != nil {
		return err
	}
	return nil
}

type TableIteratorFactory interface {
	Create(table *Table) TableIterator
}

//创建迭代器的方式用工厂方法模式
//工厂可以创建出两种具体的迭代器 randomTableIteratorFactory sortedTableIteratorFactory
type randomTableIteratorFactory struct {
}

func NewRandomTableIteratorFactory() *randomTableIteratorFactory {
	return &randomTableIteratorFactory{}
}
func (r *randomTableIteratorFactory) Create(table *Table) TableIterator {
	var records []record
	for _, r := range table.records {
		records = append(records, r)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(records), func(i, j int) {
		records[i], records[j] = records[j], records[i]
	})
	return &tableIteratorImpl{
		records: records,
		cursor:  0,
	}
}

type sortedTableIteratorFactory struct {
	Comparator Comparator
	//comparator Comparator
}

func NewSortedTableIteratorFactory(c Comparator) *sortedTableIteratorFactory {
	return &sortedTableIteratorFactory{
		Comparator: c,
	}

}

func (s *sortedTableIteratorFactory) Create(table *Table) TableIterator {
	var res []record
	for _, r := range table.records {
		res = append(res, r)
	}
	/* re := &records{
		rs:         res,
		comparator: s.Comparator,
	} */
	sort.Sort(newrecords(res, s.Comparator))
	return &tableIteratorImpl{
		records: res,
		cursor:  0,
	}
}

type records struct {
	rs         []record
	comparator Comparator
}

func newrecords(res []record, com Comparator) *records {
	return &records{
		rs:         res,
		comparator: com,
	}

}

type Comparator func(i interface{}, j interface{}) bool

//Len()
func (r *records) Len() int {
	return len(r.rs)
}

//Less(): 成绩将有低到高排序
func (r *records) Less(i, j int) bool {
	return r.comparator(r.rs[i].primaryKey, r.rs[j].primaryKey)
}

//Swap()
func (r *records) Swap(i, j int) {
	tmp := r.rs[i]
	r.rs[i] = r.rs[j]
	r.rs[j] = tmp
}
