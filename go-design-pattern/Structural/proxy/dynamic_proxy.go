package proxy

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
)

type UserInfo struct {
	ID   int
	Name string
	Age  int
}

type UserStore interface {
	// UserInfo 的具体数据获取是从具体的实现类来的
	GetUser(ctx context.Context, id int) (u *UserInfo, err error)
}

type MemoryUserStore struct {
	users []*UserInfo
}

func (m *MemoryUserStore) GetUser(ctx context.Context, id int) (u *UserInfo, err error) {
	for _, v := range m.users {
		if v.ID == id {
			return v, nil
		}
	}
	return nil, errors.New("user is not found")

}

func (m *MemoryUserStore) SetUsers(users []*UserInfo) (u *UserInfo, err error) {
	m.users = users
	return nil, errors.New("user is not found")
}

func NewMemoryUserStore() *MemoryUserStore {
	return &MemoryUserStore{}
}

// 动态增加日志 我们写一个简单的日志增强的函数，用于在方法调用后，打印这次调用的传入参数和返回值：
// 它可以对接口的实例进行动态增强，使用方法也很简单：GitHub - cocotyty/dpig: Dynamic Proxy Implementation In Go它可以对接口的实例进行动态增强
func methodLogger(in, out []reflect.Value) {
	buf := bytes.NewBuffer(nil)
	for i, value := range in {
		buf.WriteString(fmt.Sprint(value.Interface()))
		if i != len(in)-1 {
			buf.WriteString(",")
		}
	}

	// 从buf中取出
	inStr := buf.String()

	buf.Reset()
	for i, value := range out {
		buf.WriteString(fmt.Sprint(value.Interface()))
		if i != len(out)-1 {
			buf.WriteString(",")
		}
	}

	// 从buf中取出
	outStr := buf.String()
	log.Println("pass: [", inStr, "] return: [", outStr, "]")
}

func MockMemoryUserStore() *MemoryUserStore {
	return &MemoryUserStore{
		users: []*UserInfo{
			{
				ID:   1,
				Name: "Tom",
				Age:  32,
			},
			{
				ID:   2,
				Name: "Jim",
				Age:  30,
			},
			{
				ID:   3,
				Name: "Sam",
				Age:  17,
			},
		},
	}
}
