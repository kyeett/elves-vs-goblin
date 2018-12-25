package input

type Command int

const (
	MoveUp    Command = 1
	MoveDown  Command = 2
	MoveLeft  Command = 3
	MoveRight Command = 4
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
	return "Unknown"
}
