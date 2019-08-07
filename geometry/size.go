package geometry

type Dimension struct {
	Width  int
	Height int
}

func (d Dimension) Min(other Dimension) Dimension {
	return Dimension{
		Width:  Min(d.Width, other.Width),
		Height: Min(d.Height, other.Height),
	}
}

func (d Dimension) Max(other Dimension) Dimension {
	return Dimension{
		Width:  Max(d.Width, other.Width),
		Height: Max(d.Height, other.Height),
	}
}
