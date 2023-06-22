package heuristic

import (
	"fmt"

	eightpuzzle "github.com/sarnik80/8-puzzle/eightPuzzle"
)

type Manhattan struct {
	Name Lable
}

// each heuristuc has a H_value function to calculate h(node) :)
func (m Manhattan) H_value(puzzle string, goal string) int {

	var total int

	for _, number := range puzzle {

		strNumber := fmt.Sprintf("%c", number)

		total += eightpuzzle.ManhattanDistance(puzzle, goal, strNumber)

	}

	return total
}

func (m Manhattan) getName() Lable {

	return m.Name

}
