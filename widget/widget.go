package widget

type Widget interface {
	//Parent
	Constrainer
	Painter
	KeyboardHandler
}

//type Parent interface {
//	Children() []Widget
//}

type Painter interface {
	Paint(view *View)
}
