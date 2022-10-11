package structural

import (
	"context"
	"errors"
)

//
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

// 动态增加日志
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
