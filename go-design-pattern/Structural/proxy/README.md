### 代理模式的结构：

#### 1.服务接口 声明了服务接口。 代理必须遵循该接口才能伪装成服务对象。

```go
type IBusinessService interface {

    Create()

    Delete()

    Select()

    Update()

}

```

#### 2.服务 Ser­vice 类 提供了一些实用的业务逻辑。Service 实现了 IBusinessService 接口

```go
type Service struct{

}

func (s Service) Create(){

}

func (s Service) Delete(){

}

func (s Service) Select(){

}

func (s Service) Update(){

}

```

#### 3.代理 （Proxy） 类包含一个指向服务对象的引用成员变量。 代理完成其任务 （例如延迟初始化、 记录日志、 访问控制和缓存等） 后会将请求传递给服务对象。ProxyA 实现了 IBusinessService 接口 ，并组合具体的 Service

```go
// A 代理 通常情况下， 代理会对其服务对象的整个生命周期进行管理。
// 1、静态代理通常只代理一个类，动态代理是代理一个接口下的多个实现类。 2 静态代理事先知道要代理的是什么，而动态代理不知道要代理什么东西，只有在运行时才知道。
type ProxyA struct{

    realService Service

}

// 先实现服务接口
func (p *ProxyA) Create(){

     if p.CheckAccess(){

       p.realService.Create()
      }
}

func (p *ProxyA) Delete(){

     if p.CheckAccess(){

       p.realService.Delete()

     }

}

func (p *ProxyA) Select(){

   if p.CheckAccess(){

      p.realService.Select()

   }

}

func (p *ProxyA) Update(){

    if p.CheckAccess(){

       p.realService.Update()

     }
}

// 增加代理本身需要的一些操作，比如你想要给服务接口添加的操作，通过代理，添加到了代理身上

func (p *ProxyA) SetService( s Service){
      p.realService= s
}

func (p *ProxyA) CheckAccess( ){

}

```
