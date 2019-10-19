package dull

import (
	"github.com/pekim/dull/internal/geometry"
	"github.com/pekim/dull/internal/textureatlas"
)

/*
drawTextureItemToQuad adds vertices that wiil result in the gl program
blending a rectangle from the texture to a rectangle int the window.
*/
func (w *Window) drawTextureItemToQuad(
	dest geometry.RectFloat,
	textureItem *textureatlas.TextureItem,
	c Color,
) {
	tiRect := textureItem.Rect

	w.vertices = append(w.vertices,
		/*
		   1---3
		   |  /
		   | /
		   2
		*/
		dest.Left, dest.Top, tiRect.Left, tiRect.Top, c.R, c.G, c.B, c.A,
		dest.Left, dest.Bottom, tiRect.Left, tiRect.Bottom, c.R, c.G, c.B, c.A,
		dest.Right, dest.Top, tiRect.Right, tiRect.Top, c.R, c.G, c.B, c.A,

		/*
		       6
		      /|
		     / |
		   4---5
		*/
		dest.Left, dest.Bottom, tiRect.Left, tiRect.Bottom, c.R, c.G, c.B, c.A,
		dest.Right, dest.Bottom, tiRect.Right, tiRect.Bottom, c.R, c.G, c.B, c.A,
		dest.Right, dest.Top, tiRect.Right, tiRect.Top, c.R, c.G, c.B, c.A,
	)
}
