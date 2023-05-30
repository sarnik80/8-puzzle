package algorithm

import (
	"strconv"

	eightpuzzle "github.com/sarnik80/8-puzzle/eightPuzzle"
)

type algorithmName uint8

const (
	AStar = algorithmName(iota)
	Bfs
	Ids
	IDAStar
)

type Algorithm interface {
	GetName() algorithmName
	Solve(sorcePuzzle, goalPuzzle string) *eightpuzzle.EghtPuzzle
}

func (a algorithmName) String() string {
	name := []string{"AStar", "BFS", "IDS", "IDAStar"}
	i := uint8(a)
	switch {
	case i <= uint8(IDAStar):
		return name[i]
	default:
		return strconv.Itoa(int(i))
	}
}
