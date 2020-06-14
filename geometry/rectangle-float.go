package geometry

type RectFloat struct {
	Top    float64
	Bottom float64
	Left   float64
	Right  float64
}

func (r RectFloat) Width() float64 {
	return r.Right - r.Left
}

func (r RectFloat) Height() float64 {
	return r.Bottom - r.Top
}

type RectFloat32 struct {
	Top    float32
	Bottom float32
	Left   float32
	Right  float32
}

func RectFloat32From64(rect RectFloat) RectFloat32 {
	return RectFloat32{
		Top:    float32(rect.Top),
		Bottom: float32(rect.Bottom),
		Left:   float32(rect.Left),
		Right:  float32(rect.Right),
	}
}

func (r RectFloat32) Width() float32 {
	return r.Right - r.Left
}

func (r RectFloat32) Height() float32 {
	return r.Bottom - r.Top
}
