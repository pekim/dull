package widget

//func Button(r *imui.Renderer, label string, x, y float32) {
//	d := r.Drawer()
//	fg := dull.NewColor(0.5, 0.5, 0.5, 1.0)
//	bg := dull.NewColor(0.0, 0.0, 0.0, 0.0) // transparent
//
//	if r.IsFocused() {
//		bg = dull.NewColor(0.8, 0.0, 0.0, 0.3) // red
//
//		if r.KeyEvent() != nil {
//			key, _ := r.KeyEvent().Detail()
//			if key == dull.KeyTab {
//				r.FocusNext()
//			}
//		}
//	}
//
//	for i, ch := range label {
//		cell := &dull.Cell{
//			Rune: ch,
//			Fg:   fg,
//			Bg:   bg,
//		}
//		d.DrawCell(cell, x+float32(i), y)
//	}
//}
