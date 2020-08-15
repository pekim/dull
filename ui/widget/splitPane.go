package widget

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
	"github.com/pekim/dull/ui"
)

/*
	SplitPane is a widget with two children,
	and a moveable divider between them.
*/
type SplitPane struct {
	ui.BaseWidget
	orientation Orientation
	pos         int
	child1      ui.Widget
	child2      ui.Widget
}

/*
	SetOrientation sets the orientation of the split.
*/
func (sp *SplitPane) SetOrientation(orientation Orientation) {
	sp.orientation = orientation
}

/*
	SetPos sets the position of the splitter.
*/
func (sp *SplitPane) SetPos(pos int) {
	sp.pos = pos
}

/*
	SetChild1 sets the first (left or top) child.
*/
func (sp *SplitPane) SetChild1(child ui.Widget) {
	sp.child1 = child
}

/*
	SetChild2 sets the second (right or bottom) child.
*/
func (sp *SplitPane) SetChild2(child ui.Widget) {
	sp.child2 = child
}

/*
	Draw implements the Widget interface's Draw method.
*/
func (sp *SplitPane) Draw(viewport *dull.Viewport) {
	if sp.child1 != nil {
		childVp := sp.child1Viewport(viewport)
		sp.child1.Draw(childVp)
	}
	if sp.child2 != nil {
		childVp := sp.child2Viewport(viewport)
		sp.child2.Draw(childVp)
	}

	// TODO Draw splitter.
}

func (sp *SplitPane) VisitChildrenForViewport(
	viewport *dull.Viewport,
	cb ui.VisitChildViewport,
) {
	if sp.child1 != nil {
		childVp := sp.child1Viewport(viewport)
		sp.child1.VisitChildrenForViewport(childVp, cb)
	}
	if sp.child2 != nil {
		childVp := sp.child2Viewport(viewport)
		sp.child2.VisitChildrenForViewport(childVp, cb)
	}
}

func (sp *SplitPane) child1Viewport(viewport *dull.Viewport) *dull.Viewport {
	if sp.orientation == Horizontal {
		return viewport.View(geometry.RectFloat{
			Top:    0,
			Bottom: viewport.Height(),
			Left:   0,
			Right:  float64(sp.pos),
		})
	} else {
		return viewport.View(geometry.RectFloat{
			Top:    0,
			Bottom: float64(sp.pos),
			Left:   0,
			Right:  viewport.Width(),
		})
	}
}

func (sp *SplitPane) child2Viewport(viewport *dull.Viewport) *dull.Viewport {
	if sp.orientation == Horizontal {
		return viewport.View(geometry.RectFloat{
			Top:    0,
			Bottom: viewport.Height(),
			Left:   float64(sp.pos + 1),
			Right:  viewport.Width(),
		})
	} else {
		return viewport.View(geometry.RectFloat{
			Top:    float64(sp.pos + 1),
			Bottom: viewport.Height(),
			Left:   0,
			Right:  viewport.Width(),
		})
	}
}
