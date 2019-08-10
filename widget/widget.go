package widget

type Widget interface {
	//Parent
	Constrainer
	Painter
}

//type Parent interface {
//	Children() []Widget
//}

type Painter interface {
	Paint(view *View)
}
