package widget

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
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
	splitter    Border
	child1      ui.Widget
	child2      ui.Widget
}

func NewSplitPane() *SplitPane {
	splitter := Border{}
	splitter.SetEdges(EdgeLeft | EdgeRight)
	splitter.SetPosition(BorderOuter)
	splitter.SetThickness(0.2)
	splitter.SetBg(color.White)
	splitter.SetColor(color.Black)

	return &SplitPane{
		splitter: splitter,
	}
}

/*
	SetOrientation sets the orientation of the children.

	If the orientation is Horizontal, then the children will
	be arranged horizontally (left and right), and the
	splitter will run vertically.
*/
func (sp *SplitPane) SetOrientation(orientation Orientation) {
	sp.orientation = orientation

	if orientation == Horizontal {
		sp.splitter.SetEdges(EdgeLeft | EdgeRight)
	} else {
		sp.splitter.SetEdges(EdgeTop | EdgeBottom)
	}
}

/*
	SetPos sets the position of the splitter.
*/
func (sp *SplitPane) SetPos(pos int) {
	sp.pos = pos
}

func (sp *SplitPane) SetSplitterColor(color color.Color) {
	sp.splitter.SetBg(color)
}

func (sp *SplitPane) SetSplitterLineColor(color color.Color) {
	sp.splitter.SetColor(color)
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

func (sp *SplitPane) OnKey(event *dull.KeyEvent, viewport *dull.Viewport, setFocus func(widget ui.Widget)) {
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

	sp.drawSplitter(viewport)
}

func (sp *SplitPane) drawSplitter(viewport *dull.Viewport) {
	var splitterVp *dull.Viewport

	if sp.orientation == Horizontal {
		splitterVp = viewport.View(geometry.RectFloat{
			Top:    0,
			Bottom: viewport.Height(),
			Left:   float64(sp.pos),
			Right:  float64(sp.pos + 1),
		})
	} else {
		splitterVp = viewport.View(geometry.RectFloat{
			Top:    float64(sp.pos),
			Bottom: float64(sp.pos + 1),
			Left:   0,
			Right:  viewport.Width(),
		})
	}

	sp.splitter.Draw(splitterVp)
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
