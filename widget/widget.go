package widget

type Widget interface {
	Draw(v *View)
	Layout(v *View)
	PreferredSize() (int, int)
}
