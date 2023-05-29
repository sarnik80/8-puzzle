package eightpuzzle

import "strconv"

// Here we simulate C’s enum by Go’s iota for all direction

type enumDirection uint8

const (
	Up = enumDirection(iota)
	Down
	Left
	Right
)

// lable  for each direction

func (d enumDirection) String() string {
	name := []string{"up", "down", "left", "right"}
	i := uint8(d)
	switch {
	case i <= uint8(Right):
		return name[i]
	default:
		return strconv.Itoa(int(i))
	}
}

type Direction struct {
	Lable enumDirection

	X, Y int // direction coordinate
}

/*

	all directions :

	 UP   (-1 , 0)
	 Down (+1, 0)
	 Left (0, -1)
	 Right (0, +1)

*/

func defineAllDirections() *[]Direction {

	u := Direction{X: nOne, Y: zero, Lable: Up}
	d := Direction{X: one, Y: zero, Lable: Down}
	l := Direction{X: zero, Y: nOne, Lable: Left}
	r := Direction{X: zero, Y: one, Lable: Right}

	directions := []Direction{u, d, l, r}

	return &directions
}
