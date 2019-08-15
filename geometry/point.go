package geometry

type Point struct {
	X int
	Y int
}

func (p *Point) Constrain(rect Rect) {
	p.X = Max(p.X, 0)
	p.X = Min(p.X, rect.Right()-1)

	p.Y = Max(p.Y, 0)
	p.Y = Min(p.Y, rect.Bottom()-1)
}

func (p *Point) Translate(x, y int) {
	p.X += x
	p.Y += y
}
