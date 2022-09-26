package db

// demo/db/table.go

import (
	"errors"
	//"fmt"
	"math/rand"
	"reflect"
	"time"
	"strings"
)

// Table 数据表定义
type Table struct {
	name       string
	recordType reflect.Type
	records    map[interface{}]record
	// 关键点: 持有迭代器工厂方法接口
	iteratorFactory TableIteratorFactory
}

func NewTable(name string) *Table {
	return &Table{
		name:    name,
		records: make(map[interface{}]record),
		iteratorFactory: NewRandomTableIteratorFactory(),
	}
}


func (t *Table) WithType(recordType reflect.Type) *Table {
	t.recordType = recordType
	return t
}

func (t *Table) Name() string {
	return strings.ToLower(t.name)
}

func (t *Table) QueryByPrimaryKey(key interface{}, value interface{}) error {
	record, ok := t.records[key]
	if !ok {
		return ErrRecordNotFound
	}
	return record.convertByValue(value)
}

func (t *Table) Insert(key interface{}, value interface{}) error {

	if _, ok := t.records[key]; ok {
		return errors.New("ErrPrimaryKeyConflict")
	}
	record, err := recordFrom(key, value)
	if err != nil {
		return err
	}
	t.records[key] = record
	return nil

}

func (t *Table) Update(key interface{}, value interface{}) error {
	if _, ok := t.records[key]; !ok {
		return ErrRecordNotFound
	}
	record, err := recordFrom(key, value)
	if err != nil {
		return err
	}
	t.records[key] = record
	return nil
}

func (t *Table) Delete(key interface{}) error {
	if _, ok := t.records[key]; !ok {
		return ErrRecordNotFound
	}
	delete(t.records, key)
	return nil
}

// 关键点: 定义Setter方法，提供迭代器工厂的依赖注入
func (t *Table) WithTableIteratorFactory(iteratorFactory TableIteratorFactory) *Table {
	t.iteratorFactory = iteratorFactory
	return t
}

// 关键点: 定义创建迭代器的接口，其中调用迭代器工厂完成实例化
func (t *Table) Iterator() TableIterator {
	return t.iteratorFactory.Create(t)
}

type Next func(interface{}) error
type HasNext func() bool

// 迭代器模式 go风格的实现
// Go 风格的实现，利用了函数闭包的特点，把原本在迭代器实现的逻辑，放到了迭代器创建方法上。
// 相比面向对象风格，省掉了迭代器抽象接口和实现对象的定义，看起来更加的简洁
// 声明 HashNext 和 Next 的函数类型，等同于迭代器抽象接口的作用
func (t *Table) ClosureIterator() (HasNext, Next) {
	var records []record
	for _, r := range t.records {
		records = append(records, r)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(records), func(i, j int) {
		records[i], records[j] = records[j], records[i]
	})
	size := len(records)
	cursor := 0
	hasNext := func() bool {
		return cursor < size
	}
	next := func(next interface{}) error {
		record := records[cursor]
		cursor++
		if err := record.convertByValue(next); err != nil {
			return err
		}
		return nil
	}
	return hasNext, next
}

// 迭代器模式 用channel 来实现
