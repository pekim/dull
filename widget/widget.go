package widget

type Widget interface {
	Draw(v *View)
	Layout(v *View)
	PreferredSize(v *View) (int, int)
}
