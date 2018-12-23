package actions

import (
	"github.com/kyeett/elves-vs-goblin/pkg/geom"
)

type Action int

const (
	Move Action = iota + 1
	_
	_
	Build
)

type Signal struct {
	ID     string
	Coord  geom.Coord
	Action Action
}
