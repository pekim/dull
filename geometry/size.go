package geometry

type Size struct {
	Width  int
	Height int
}

func (d Size) Min(other Size) Size {
	return Size{
		Width:  Min(d.Width, other.Width),
		Height: Min(d.Height, other.Height),
	}
}

func (d Size) Max(other Size) Size {
	return Size{
		Width:  Max(d.Width, other.Width),
		Height: Max(d.Height, other.Height),
	}
}
