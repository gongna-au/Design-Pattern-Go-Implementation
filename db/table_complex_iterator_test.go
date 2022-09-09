package db

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTableComplexIterator(t *testing.T) {
	iteratorFactory := NewcomplexTableIteratorFactory()
	// 关键点5: 使用时，直接通过for-range来遍历channel读取记录
	table := NewTable("testRegion").WithType(reflect.TypeOf(new(testRegion))).
		WithTableIteratorFactory(NewSortedTableIteratorFactory(regionIdLess))
	table.Insert(3, &testRegion{Id: 3, Name: "beijing"})
	table.Insert(1, &testRegion{Id: 1, Name: "shanghai"})
	table.Insert(2, &testRegion{Id: 2, Name: "guangdong"})

	iterator := iteratorFactory.Create(table)
	if v, ok := iterator.(*complextableIteratorImpl); ok {
		for record := range v.Iterator() {
			fmt.Println(record)
		}
	}
}
