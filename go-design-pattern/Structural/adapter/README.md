### 对象适配器模式的结构

> 新增的适配器实现对外的接口，继承第三方类

#### 1. IClient 一个面向客户端的接口，客户端必须和这个接口龚作

```go
// 虚线+空三角形 实现
// 实线+空三角形 继承
// 实线+空平行四边形 组合接口 组合结构体
// 客户端必须通过客户端接口和适配器工作
// 适配器类向上实现了对外提供的接口，向下继承第三方类，实现对外提供的接口的本质是调用被继承类的已有的方法
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

func (a *Adapter) Update(i int) {
	a.sd.UpdateUser(i)
}

```

#### 3. 一个具体的服务，作为参数传递给 Adapter

```go
// 具体的服务
type ServiceUpdate struct {
}

func (s *ServiceUpdate) UpdateUser(user int) {

}

func NewServiceUpdate()*ServiceUpdate {
    return &ServiceUpdate{

    }
}

func Client() {
	adapter := NewAdapter(NewServiceUpdate())
	adapter.Update(9)
}
```

#### 4. 外界通过 Adapter 实现的接口 IClient 和 Adapter 进行交互

### 类适配器模式的结构

#### 1. RoundHole 一个已经存在的类

```go
// 圆柱体
type RoundHole struct {
	Radius int
}

func NewRoundHole(radius int) *RoundHole {
	return &RoundHole{
		Radius: radius,
	}
}

func (e *RoundHole) GetRadius() int {
	return e.Radius
}

func (e *RoundHole) Fit(r *RoundPeg) bool {
	return e.GetRadius() == r.GetRadius()
}

```

#### 2. RoundPeg 另外一个已经存在的类

```go
// 圆柱体
type RoundPeg struct {
	Radius int
}

func NewRoundPeg(radius int) *RoundPeg {
	return &RoundPeg{
		Radius: radius,
	}
}

func (s *RoundPeg) GetRadius() int {
	return s.Radius
}

```

#### 3. SquarePeg 另外一个存在的类

```go
// 长方体
type SquarePeg struct {
	Width int
}

func NewSquarePeg(width int) *SquarePeg {
	return &SquarePeg{
		Width: width,
	}
}

func (s *SquarePeg) GetWidth() int {
	return s.Width
}

```

#### 4. SquarePegAdapter 一个适配器，使得 SquarePeg 适应 RoundPeg 的某个方法 ，继承 RoundPeg 使得 SquarePeg 适配 RoundPeg，组合 SquarePeg 实现 RoundPeg 的 GetRadius()方法，使得 SquarePeg 看起来也可以获取到半径

```go
// 适配器有两个方向的，使得谁适配谁很重要，A 被拿来适配 B ，那么适配器要和A合作，要继承B,
type SquarePegAdapter struct {
	// 继承 RoundPeg 才能使得 SquarePeg 适配 RoundPeg
	*RoundPeg
	// 保存SquarePeg
	peg *SquarePeg
}

func NewSquarePegAdapter(r *RoundPeg) *SquarePegAdapter {
	return &SquarePegAdapter{
		RoundPeg: r,
	}
}

// SquarePeg做了参数
func (s *SquarePegAdapter) Adapter(r *SquarePeg) {
	s.peg = r
}

// 实现了 GetRadius()
func (s *SquarePegAdapter) GetRadius() int {
	return s.peg.GetWidth() * int(math.Sqrt(2))
}
```

#### 5. 客户端是如何做的？显然，是因为客户端（外界）不可以和 SquarePeg 直接交互，只能和 RoundPeg 进行交互 ，所以适配器继承 RoundPeg 并保存、操作 SquarePeg 使得 SquarePeg 看起来和 RoundPeg 一样，也能获取到 GetRadius()

```go


```
