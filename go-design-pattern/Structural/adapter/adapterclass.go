package adapter

import "math"

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

func (e *RoundHole) Fit(r IRoundPeg) bool {
	return e.GetRadius() == r.GetRadius()
}

type IRoundPeg interface {
	GetRadius() int
}

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

// 适配器有两个方向的，使得谁适配谁很重要，A 被拿来适配 B ，那么适配器要和A合作，要继承B,
type SquarePegAdapter struct {
	// 继承 RoundPeg 才能使得 SquarePeg 适配 RoundPeg
	*RoundPeg
	// 保存SquarePeg
	peg *SquarePeg
}

func NewSquarePegAdapter() *SquarePegAdapter {
	return &SquarePegAdapter{}
}

// SquarePeg做了参数
func (s *SquarePegAdapter) Adapter(r *SquarePeg) *SquarePegAdapter {
	s.peg = r
	return s
}

// 实现了 GetRadius()
func (s *SquarePegAdapter) GetRadius() int {
	return s.peg.GetWidth() * int(math.Sqrt(2))
}

func GetRadiusClient() {
	// 为什么要为父类声明接口类型
	// 子类指针不可以赋值给父类指针
	var R IRoundPeg
	round := NewRoundPeg(3)
	square := NewSquarePeg(6)
	R = round
	R.GetRadius()
	// 得到适配器，把要被适配的square 传进去
	adapter := NewSquarePegAdapter().Adapter(square)
	// 把适配器赋值给R
	R = adapter
	R.GetRadius()

}
