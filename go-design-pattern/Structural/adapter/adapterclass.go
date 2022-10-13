package adapter

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

// 圆柱体
type RoundPeg struct {
	Radius int
}

func NewRoundPeg(radius int) *RoundPeg {
	return &RoundPeg{
		Radius: radius,
	}
}

func (s *RoundPeg) GetWidth() int {
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
}

func NewSquarePegAdapter(r *RoundPeg) *SquarePegAdapter {
	return &SquarePegAdapter{
		r,
	}

}

// SquarePeg做了参数
func (s *SquarePegAdapter) Adapter(r *SquarePeg) {

}
