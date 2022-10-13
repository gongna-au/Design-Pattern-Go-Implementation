package composite

import "fmt"

type Graphic interface {
	Move(x int, y int)
	Draw()
}

type Dot struct {
	x int
	y int
}

func (d *Dot) Move(x int, y int) {
	d.x = d.x + x
	d.y = d.y + y
}
func NewDot(x int, y int) *Dot {
	return &Dot{
		x: x,
		y: y,
	}
}

func (d *Dot) Draw() {
	fmt.Println("Draw over")
}

type Circle struct {
	*Dot
	radius int
}

func NewCircle(d *Dot, radius int) *Circle {
	return &Circle{
		Dot:    d,
		radius: radius,
	}
}

func (c *Circle) Move(x int, y int) {
	c.x = c.x + x
	c.y = c.y + y
}

func (c *Circle) Draw() {
	fmt.Println("Circle draw over")
}

type CompoundGraphic struct {
	graphic []Graphic
}

func NewCompoundGraphic() *CompoundGraphic {
	return &CompoundGraphic{
		[]Graphic{},
	}
}

func (c CompoundGraphic) Move(x int, y int) {
	for _, v := range c.graphic {
		v.Move(x, y)
	}
}

func (c CompoundGraphic) Draw() {
	for _, v := range c.graphic {
		v.Draw()
	}
}

func (c CompoundGraphic) AddGraphics(g ...Graphic) {
	c.graphic = append(c.graphic, g...)
}

func (c CompoundGraphic) GetGraphics(i int) Graphic {
	return c.graphic[i]
}

func GraphicClient() {
	var c Graphic
	temp := NewCompoundGraphic()
	temp.AddGraphics(NewCircle(NewDot(1, 2), 3), NewCircle(NewDot(4, 5), 6))
	c = temp
	c.Draw()
	c.Move(1, 1)
}
