package world

import (
	"bytes"

	"github.com/kyeett/elves-vs-goblin/pkg/geom"
)

// The World contains a map and players
type World struct {
	m    [][]byte
	Size geom.Rect
}

const size = 5

// NewDefaultWorld returns a 9x9 world filled with empty spaces
func NewDefaultWorld() World {
	m := make([][]byte, size)
	for y := 0; y < size; y++ {
		m[y] = bytes.Repeat([]byte("-"), size)
	}

	return World{
		m: m,
		Size: geom.Rect{
			W: size,
			H: size,
		},
	}
}

// Center returns the center position of the world, zero indexed
func (w World) Center() geom.Coord {
	return geom.Coord{
		X: len(w.m[0]) / 2,
		Y: len(w.m) / 2,
	}
}

// Dims returns a rectangle struct representing the width and height of the world
func (w World) Dims() geom.Rect {
	return geom.Rect{
		W: len(w.m[0]),
		H: len(w.m),
	}
}

func (w World) String() string {
	var buffer bytes.Buffer
	for _, row := range w.m {
		buffer.Write(row)
		buffer.WriteString("\n")
	}
	return buffer.String()
}

func (w World) Rows() [][]byte {
	return w.m
}

// const paddingX = 3

// var horizontalLine = strings.Repeat("-", paddingX*width)

// func (w World) String() string {
// 	sY, sX := w.Dims()

// 	grid := make([][]rune, sY)
// 	s := ""
// 	for y := 0; y < sY; y++ {
// 		grid[y] = make([]rune, paddingX*sX)
// 		for x := 0; x < sX*paddingX; x += paddingX {

// 			switch w.At(y, x/paddingX) {
// 			case Tree:
// 				grid[y][x] = '🌲'
// 			case RoundedTree:
// 				grid[y][x] = '🌳'

// 			case Whale:
// 				grid[y][x] = '🐳'
// 			default:
// 				grid[y][x] = ' '

// 			}

// 			grid[y][x+1] = ' '
// 			grid[y][x+2] = '|'
// 		}
// 		s += string(grid[y]) + "\n"
// 		s += horizontalLine + "\n"
// 	}

// 	return s
// }
