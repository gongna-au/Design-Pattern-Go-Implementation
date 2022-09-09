package db

import (
	//"strings"
	"sync"
)

var memoryDbInstance = &memoryDb{
	tables: sync.Map{},
}

func MemoryDbInstance() *memoryDb {
	return memoryDbInstance
}

// memoryDb 内存数据库
type memoryDb struct {
	tables sync.Map // key为tableName，value为table
}

func (m *memoryDb) CreateTable(t *Table) {

}
