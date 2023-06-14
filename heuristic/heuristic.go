package heuristic

import (
	"strconv"

	eightpuzzle "github.com/sarnik80/8-puzzle/eightPuzzle"
)

type Heuristic interface {
	getName() Lable
	H_value(puzzle string, goal string) int
}

type Lable uint8

const (
	ManhattanDistance = Lable(iota)
)

func (d Lable) String() string {
	name := []string{"ManhattanDistance"}
	i := uint8(d)
	switch {
	case i <= uint8(eightpuzzle.Right):
		return name[i]
	default:
		return strconv.Itoa(int(i))
	}
}
