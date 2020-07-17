package widget

/*
	Edge represents one or more edges (sides) of
	a rectangle.

	Edges can be bitwise ORed.
	For example EdgeTop | EdgeBottom represents
	both the top and bottom edges.
*/
type Edge int

const (
	// EdgeTop is  the top edge.
	EdgeTop Edge = 1 << iota
	// EdgeBottom is the bottom edge.
	EdgeBottom
	// EdgeLeft is the left edge.
	EdgeLeft
	// EdgeRight is the right edge.
	EdgeRight

	// EdgeAll represents all four edges.
	EdgeAll = EdgeTop | EdgeBottom | EdgeLeft | EdgeRight
)
