package adapter

// 适配器实现具体对象的接口，包装具体对象，然后外界通过接口和适配器进行沟通
// 如果服务类的接口是可能被更改的，那么写一个不会发生变化的接口，适配器实现了这个接口。“客户端代码通过客户端接口与适配器一起工作“
// 实现了增加新的适配器来达到
// 暴露给客户端的接口
type IClient interface {
	Update(int)
}

// Adapter实现了接口，供给外界的client 使用
type Adapter struct {
	sd *ServiceUpdate
}

func NewAdapter(sd *ServiceUpdate) *Adapter {
	return &Adapter{
		sd: sd,
	}
}

func (a *Adapter) Update(i int) {
	a.sd.UpdateUser(i)
}

// 具体的服务有不同与接口的方法
type ServiceUpdate struct {
}

func (s *ServiceUpdate) UpdateUser(user int) {

}

// 作为参数传递给Adapter
func NewServiceUpdate() *ServiceUpdate {
	return &ServiceUpdate{}
}

func Client() {
	adapter := NewAdapter(NewServiceUpdate())
	adapter.Update(9)
}
