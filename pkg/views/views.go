package views

import (
	"bytes"
	"io"

	"github.com/kyeett/elves-vs-goblin/pkg/world"

	"github.com/kyeett/elves-vs-goblin/pkg/geom"
)

type View struct {
	w       io.Writer
	padding []byte
	size    geom.Rect
}

const size = 9

func New(w io.Writer) View {
	return View{
		w:       w,
		padding: []byte("~"),
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

func (v View) Draw(w *world.World) {
	var buffer bytes.Buffer

	missing := v.size.Sub(w.Size)
	row, before, after := paddingBytes(v, missing)
	for y := 0; y < missing.H/2; y++ {
		buffer.Write(row)
	}

	for _, r := range w.Rows() {
		buffer.Write(before)
		buffer.Write(r)
		// v.w.Write(r)

		buffer.Write(after)
		buffer.WriteString("\n")
	}

	for y := 0; y < missing.H-missing.H/2; y++ {
		buffer.Write(row)
	}

	// v.w.Write([]byte("OK"))
	// v.w.Write(buffer.Bytes())
	buffer.WriteTo(v.w)
	v.w.Write([]byte(""))

	// Todo why is this needed?
}
