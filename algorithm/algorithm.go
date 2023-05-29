package algorithm

import "strconv"

type algorithm uint8

const (
	AStar = algorithm(iota)
	BFS
	IDS
	IDAStar
)

type Algorithm interface {
	getName()
	solve()
}

func (a algorithm) String() string {
	name := []string{"AStar", "BFS", "IDS", "IDAStar"}
	i := uint8(a)
	switch {
	case i <= uint8(IDAStar):
		return name[i]
	default:
		return strconv.Itoa(int(i))
	}
}
