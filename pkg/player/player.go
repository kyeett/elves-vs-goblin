package player

import (
	"crypto/md5"
	"encoding/hex"
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

func (p *Player) Goto(x, y int) {
	p.X = x
	p.Y = y
}

func NewPlayer() Player {
	hash := md5.New()
	hash.Write([]byte(strconv.Itoa(rand.Intn(123456))))
	ID := hex.EncodeToString(hash.Sum(nil))[0:8]
	return Player{
		ID,
		geom.Coord{0, 0},
	}

}
