package db

import (
	"fmt"
	"reflect"
	"testing"
)

type testRegion struct {
	Id   int
	Name string
}

func TestTable(t *testing.T) {
	tableName := "testRegion"
	table := NewTable(tableName).WithType((reflect.TypeOf(new(testRegion))))
	table.Insert(2, &testRegion{Id: 2, Name: "beijing"})
	fmt.Println(table.records)

}
