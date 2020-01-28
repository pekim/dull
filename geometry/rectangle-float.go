package geometry

type RectFloat struct {
	Top    float32
	Bottom float32
	Left   float32
	Right  float32
}

func (r RectFloat) Width() float32 {
	return r.Right - r.Left
}

func (r RectFloat) Height() float32 {
	return r.Bottom - r.Top
}
