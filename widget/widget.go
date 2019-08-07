package widget

type Widget interface {
	//Paint(v *View)
	//Layout(v *View)
	//PreferredSize() (int, int)

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
