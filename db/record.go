package db

import (
	"errors"
	//"fmt"
	"reflect"
	"strings"
)

// 因为数据库的每个表都存储着不同对象
// 所以需要把类型存进去，根据类型创建自己需要的对象，再根据对象的属性，创建出表的每一列的属性
// 其中，Table 底层用 map 存储对象数据，但并没有存储对象本身，而是从对象转换而成的 record
type record struct {
	primaryKey interface{}
	fields     map[string]int
	values     []interface{}
}

//从对象转化为 record
func recordFrom(key interface{}, value interface{}) (r record, e error) {
	defer func() {
		if err := recover(); err != nil {
			r = record{}
			e = errors.New("ErrRecordTypeInvalid")
		}
	}()

	vType := reflect.TypeOf(value)
	//fmt.Println("vType:", vType)
	vVal := reflect.ValueOf(value)
	//fmt.Println("vVal:", vVal)

	if vVal.Type().Kind() == reflect.Ptr {
		//fmt.Println("is ptr")
		vType = vType.Elem()
		//fmt.Println("vType:", vType)

		vVal = vVal.Elem()
		//fmt.Println("vVal:", vVal)

	}

	record := record{
		primaryKey: key,
		fields:     make(map[string]int, vVal.NumField()),
		values:     make([]interface{}, vVal.NumField()),
	}
	//fmt.Println("vVal.NumField()", vVal.NumField())

	for i := 0; i < vVal.NumField(); i++ {
		fieldType := vType.Field(i)
		//fmt.Println("fieldType :", fieldType)
		fieldVal := vVal.Field(i)
		//fmt.Println("fieldVal:", fieldVal)
		name := strings.ToLower(fieldType.Name)
		record.fields[name] = i
		record.values[i] = fieldVal.Interface()
	}

	return record, nil

}

func (r record) convertByValue(result interface{}) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = errors.New("ErrRecordTypeInvalid")
		}
	}()
	rType := reflect.TypeOf(result)
	rVal := reflect.ValueOf(result)
	if rType.Kind() == reflect.Ptr {
		rType = rType.Elem()
		rVal = rVal.Elem()
	}
	for i := 0; i < rType.NumField(); i++ {
		field := rVal.Field(i)
		field.Set(reflect.ValueOf(r.values[i]))
	}
	return nil
}
