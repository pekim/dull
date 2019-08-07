package geometry

type Point struct {
	X int
	Y int
}

func (p Point) Translate(x, y int) {
	p.X += x
	p.Y += y
}
