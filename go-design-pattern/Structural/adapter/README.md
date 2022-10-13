### 对象适配器模式的结构

#### 1. IClient 一个面向客户端的接口，客户端必须和这个接口龚作

```go
// 虚线+空三角形 实现
// 实线+空三角形 继承
// 实线+空平行四边形 组合接口 组合结构体
// 客户端必须通过客户端接口和适配器工作
type IClient interface {
	Update(int)
}

```

#### 2. 一个具体的适配器，必须实现 IClient 接口，且和要被适配的对象一起工作

```go
type Adapter struct {
	sd *ServiceUpdate
}

func NewAdapter(sd *ServiceUpdate) *Adapter {
	return &Adapter{
		sd: sd,
	}
}

func (a *Adapter) Update(int) {
	a.sd.UpdateUser(1)
}

```

#### 3. 一个具体的适配器，必须实现 IClient 接口，且和要被适配的对象一起工作

```go
// 具体的服务实现了接口的方法
type ServiceUpdate struct {
}

func (s *ServiceUpdate) UpdateUser(user int) {

}
func NewServiceUpdate()*ServiceUpdate {
    return &ServiceUpdate{

    }
}
```

### 类适配器模式的结构

#### 1. IClient 一个面向客户端的接口，客户端必须和这个接口龚作

```go
// 虚线+空三角形 实现
// 实线+空三角形 继承
// 实线+空平行四边形 组合接口 组合结构体
// 实线 做参数
// 客户端必须通过客户端接口和适配器工作
type IClient interface {
	Update(int)
}

```
