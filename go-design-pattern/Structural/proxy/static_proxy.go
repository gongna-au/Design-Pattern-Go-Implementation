package proxy

import (
	"fmt"
	"log"
	"time"
)

// 静态代理和动态代理的区别
// IUser IUser
type IUser interface {
	Login(username, password string) error
}

// User 用户
// @proxy IUser
type User struct {
}

// Login 用户登录
func (u *User) Login(username, password string) error {
	// 不实现细节
	fmt.Println("User login")
	return nil
}

type UserProxy struct {
	user *User
}

// Login 用户登录
func (u *UserProxy) SetUser(user *User) error {
	// 不实现细节
	u.user = user
	fmt.Println("UserProxy SetUser")
	return nil
}

// Login 代理底层调用user.Login()登陆
func (u *UserProxy) Login(username, password string) error {
	// before 这里可能会有一些统计的逻辑
	start := time.Now()
	// 这里是原有的业务逻辑
	if err := u.user.Login(username, password); err != nil {
		return err
	}
	// after 这里可能也有一些监控统计的逻辑
	log.Printf("user login cost time: %s", time.Now().Sub(start))
	return nil
}

// 代理做一些额外的操作，比如检查
func (u *UserProxy) Check(username, password string) bool {
	if username == "" || password == "" {
		return false
	}
	if len(password) < 3 {
		return false
	}
	return true
}

func NewUser() *User {
	return &User{}
}

// 提供给客户端，让客户端得到代理，然后操作代理来实现功能
func NewUserProxy() *UserProxy {
	return &UserProxy{}
}
