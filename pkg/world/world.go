package world

import (
	"bytes"
	"fmt"
	"sync"
	"time"

	"github.com/kyeett/elves-vs-goblin/pkg/player"

	"github.com/kyeett/elves-vs-goblin/pkg/geom"
)

// The World contains a map and players
type World struct {
	m       [][]byte
	Size    geom.Rect
	players []*player.Player
	mut     *sync.RWMutex
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
		mut: &sync.RWMutex{},
	}
}

// Center returns the center position of the world, zero indexed
func (w World) Center() geom.Coord {
	return geom.Coord{
		X: len(w.m[0]) / 2,
		Y: len(w.m) / 2,
	}
}

// NewPlayer adds a player to the world
func (w *World) NewPlayer(p *player.Player) {
	w.mut.Lock()
	w.players = append(w.players, p)
	w.mut.Unlock()
}

func (w World) Rows() [][]byte {
	w.mut.RLock()
	duplicate := make([][]byte, len(w.m))
	for i := range w.m {
		duplicate[i] = make([]byte, len(w.m[i]))
		copy(duplicate[i], w.m[i])
	}

	for _, p := range w.players {
		fmt.Println("Yay")
		duplicate[p.Y][p.X] = '#'
	}
	w.mut.RUnlock()
	return duplicate
}

func (w *World) Start() {
	p := player.NewPlayer()
	w.NewPlayer(&p)

	for {
		// Get user input
		p.Move(1, 0)
		time.Sleep(200 * time.Millisecond)
		p.Move(0, 1)
		time.Sleep(200 * time.Millisecond)
		p.Move(0, 1)
		time.Sleep(200 * time.Millisecond)
		p.Move(1, 0)
		time.Sleep(200 * time.Millisecond)

		p.Move(0, -1)
		time.Sleep(200 * time.Millisecond)
		p.Move(-1, 0)
		time.Sleep(200 * time.Millisecond)
		p.Move(0, -1)
		time.Sleep(200 * time.Millisecond)
		p.Move(-1, 0)
		time.Sleep(200 * time.Millisecond)

		// Update game state
	}
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
