package views

import (
	"bytes"

	"github.com/kyeett/elves-vs-goblin/pkg/geom"
	"github.com/kyeett/elves-vs-goblin/pkg/world"
)

type View struct {
	world   *world.World
	padding []byte
	center  geom.Coord
	size    geom.Rect
}

const size = 9

func NewView(w *world.World) View {
	return View{
		world:   w,
		padding: []byte("."),
		center:  w.Center(),
		size: geom.Rect{
			W: size,
			H: size,
		},
	}
}

func paddingBytes(v View, missing geom.Rect) ([]byte, []byte, []byte) {
	row := bytes.Repeat(v.padding, v.size.W)
	row = append(row, []byte("\n")...)
	before := bytes.Repeat(v.padding, missing.W/2)
	after := bytes.Repeat(v.padding, missing.W-missing.W/2)
	return row, before, after
}

func (v View) String() string {
	var buffer bytes.Buffer
	w := *v.world

	missing := v.size.Sub(w.Size)

	row, before, after := paddingBytes(v, missing)
	for y := 0; y < missing.H/2; y++ {
		buffer.Write(row)
	}

	for _, r := range w.Rows() {
		buffer.Write(before)
		buffer.Write(r)
		buffer.Write(after)
		buffer.WriteString("\n")
	}

	for y := 0; y < missing.H-missing.H/2; y++ {
		buffer.Write(row)
	}

	return buffer.String()
}
