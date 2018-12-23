package world

import (
	"bytes"
	"sync"
	"time"

	"github.com/kyeett/elves-vs-goblin/pkg/transport"

	"github.com/kyeett/elves-vs-goblin/pkg/player"

	"github.com/kyeett/elves-vs-goblin/pkg/geom"
)

// The World contains a map and players
type World struct {
	m       [][]byte
	Size    geom.Rect
	players []*player.Player
	msger   transport.Messager
	mut     *sync.RWMutex
}

const size = 5

// NewDefaultWorld returns a 5x5 world filled with empty spaces
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

func NewWorld() World {
	w := NewDefaultWorld()
	w.msger = transport.DefaultNats()
	return w
}

// Center returns the center position of the world, zero indexed
func (w World) Center() geom.Coord {
	return geom.Coord{
		X: len(w.m[0]) / 2,
		Y: len(w.m) / 2,
	}
}

// NewPlayer adds a player to the world
func (w *World) AddPlayer(p *player.Player) {
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
		duplicate[p.Y][p.X] = '#'
	}
	w.mut.RUnlock()
	return duplicate
}

func (w *World) Start() {
	p := player.NewDefaultPlayer()
	w.AddPlayer(&p)

	for {
		// Get user input
		time.Sleep(200 * time.Millisecond)

		time.Sleep(200 * time.Millisecond)
		time.Sleep(200 * time.Millisecond)
		time.Sleep(200 * time.Millisecond)

		time.Sleep(200 * time.Millisecond)
		time.Sleep(200 * time.Millisecond)
		time.Sleep(200 * time.Millisecond)
		time.Sleep(200 * time.Millisecond)
		// c.Publish(, v interface{})

		// Update game state
	}
}

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
