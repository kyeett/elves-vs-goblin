package world

import (
	"bytes"
	"encoding/json"
	"sync"

	"github.com/kyeett/elves-vs-goblin/pkg/player"

	"github.com/kyeett/elves-vs-goblin/pkg/geom"
)

// The World contains a map and players
type World struct {
	M       [][]byte
	Size    geom.Rect
	Players []*player.Player
}

var mut = sync.RWMutex{}

const size = 5

// NewDefaultWorld returns a 5x5 world filled with empty spaces
func NewDefaultWorld() World {
	m := make([][]byte, size)
	for y := 0; y < size; y++ {
		m[y] = bytes.Repeat([]byte(" "), size)
	}

	return World{
		M: m,
		Size: geom.Rect{
			W: size,
			H: size,
		},
	}
}

// Center returns the center position of the world, zero indexed
func (w World) Center() geom.Coord {
	return geom.Coord{
		X: len(w.M[0]) / 2,
		Y: len(w.M) / 2,
	}
}

// NewPlayer adds a player to the world
func (w *World) AddPlayer(p *player.Player) {
	mut.Lock()
	w.Players = append(w.Players, p)
	mut.Unlock()
}

func (w World) Rows() [][]byte {
	mut.RLock()
	duplicate := make([][]byte, len(w.M))
	for i := range w.M {
		duplicate[i] = make([]byte, len(w.M[i]))
		copy(duplicate[i], w.M[i])
	}

	for _, p := range w.Players {
		duplicate[p.Y][p.X] = '#'
	}
	mut.RUnlock()
	return duplicate
}

func (w *World) Start() {
	p := player.NewDefaultPlayer()
	w.AddPlayer(&p)

}

// Todo: not a good solution, ask Slack
func (w *World) UnmarshalJSON(data []byte) error {

	aux := &struct {
		M       [][]byte
		Size    geom.Rect
		Players []player.Player
	}{}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var players []*player.Player
	for _, p := range aux.Players {
		players = append(players, &p)
	}

	w.M = aux.M
	w.Size = aux.Size
	w.Players = players
	return nil
}
