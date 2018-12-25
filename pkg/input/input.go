package input

import "fmt"

type Command int

const (
	MoveUp    Command = 0
	MoveDown  Command = 1
	MoveLeft  Command = 2
	MoveRight Command = 3
)

func (c Command) String() string {
	switch c {
	case MoveUp:
		return "MoveUp"
	case MoveDown:
		return "MoveDown"
	case MoveLeft:
		return "MoveLeft"
	case MoveRight:
		return "MoveRight"
	}
	return fmt.Sprintf("Unknown %d", c)
}
