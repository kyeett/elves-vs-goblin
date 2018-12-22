package player

import (
	"crypto/md5"
	"math/rand"
	"strconv"

	"github.com/kyeett/elves-vs-goblin/pkg/geom"
)

type Player struct {
	ID string
	geom.Coord
}

func (p *Player) Move(x, y int) {
	p.X += x
	p.Y += y
}

func NewPlayer() Player {
	id := md5.Sum([]byte(strconv.Itoa(rand.Intn(123456))))
	return Player{
		string(id[:10]),
		geom.Coord{0, 0},
	}

}
