package widget

type Widget interface {
	Parent
	Constrainer
	Painter
	KeyboardHandler
}

type Parent interface {
	Children() []Widget
}

type Childless struct{}

func (c Childless) Children() []Widget {
	return []Widget{}
}

type Painter interface {
	Paint(view *View, root *Root)
}
