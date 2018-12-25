package player

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/kyeett/elves-vs-goblin/pkg/geom"
)

type Player struct {
	ID string
	geom.Coord
}

func (p *Player) Goto(x, y int) {
	p.X = x
	p.Y = y
}

func NewDefaultPlayer() Player {
	hash := md5.New()
	hash.Write([]byte(strconv.Itoa(rand.Intn(123456))))
	ID := hex.EncodeToString(hash.Sum(nil))[0:8]

	return NewPlayer(ID)
}

func NewPlayer(ID string) Player {
	return Player{
		ID,
		geom.Coord{0, 0},
	}
}

func (p Player) String() string {
	return fmt.Sprintf("[ID:%s %s]", p.ID, p.Coord)
}

// func (p *Player) MarshalJSON() ([]byte, error) {
// 	b, err := json.Marshal(*p)
// 	if err != nil {
// 		return nil, err
// 	}
// 	log.Error("Marshal")
// 	return b, nil
// }
