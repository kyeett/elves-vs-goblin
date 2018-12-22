package geom

// Coord holds x and y coordinates
type Coord struct {
	X, Y int
}

// Rect holds width and height
type Rect struct {
	W, H int
}

// Sub returns the differences in width and height between two rectangles
func (r Rect) Sub(s Rect) Rect {
	return Rect{
		W: r.W - s.W,
		H: r.H - s.H,
	}
}
