package geom

import "fmt"

// Coord holds x and y coordinates
type Coord struct {
	X, Y int
}

// Add returns a new coordinate by offset argument
func (c Coord) Add(dx, dy int) Coord {
	return Coord{
		X: c.X + dx,
		Y: c.Y + dy,
	}
}

func (c Coord) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}

// Rect holds width and height
type Rect struct {
	W, H int
}

func (r Rect) String() string {
	return fmt.Sprintf("(%d,%d)", r.W, r.H)
}

// Sub returns the differences in width and height between two rectangles
func (r Rect) Sub(s Rect) Rect {
	return Rect{
		W: r.W - s.W,
		H: r.H - s.H,
	}
}
