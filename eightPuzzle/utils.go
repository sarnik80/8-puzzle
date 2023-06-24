package eightpuzzle

import (
	"math"
	"strings"
)

/*
	**For easier work, I considered my puzzle as a string


	8   4   3
	7   0   1   => puzzle
	6   5   2


	123456780   => string


	index (0)  =>  in string is  :  4


*/

func FindIndex(puzzle, numbrt string) (i, j int) {

	index := strings.Index(puzzle, numbrt)

	i = index / MaxRow

	j = index - (i * MaxCol)

	return

}

//  Calculates the coordinates of a number in a puzzle

func getCoordinates(puzzle, strNumber string) (x, y int) {

	x, y = FindIndex(puzzle, strNumber)

	return

}

func findAvailableDirections(i, j int) *[]Direction {

	// a slice of directions
	directions := *defineAllDirections()
	avDirections := []Direction{}

	for _, dir := range directions {

		if !(i+dir.X < 0) && !(j+dir.Y > 2) && !(j+dir.Y < 0) && !(i+dir.X > 2) {

			avDirections = append(avDirections, dir)
		}

	}

	return &avDirections

}

// calculates sum of two index

func calCulateDistIndex(source, dist int) int {

	return source + dist
}

//  convert "123456780"  to   [1,2,3,4,5,6,7,8,0]

func slicedString(data string) []string {
	result := strings.Split(data, "")

	return result
}

//To convert a slice to a string  [1,2,3,4,5,6,7,8,0]  to  "123456780"
func stringedSlice(strSlice []string) string {
	return strings.Join(strSlice, "")
}

/*

 By having the coordinates of a tile in the
 current puzzle and in the target puzzle,
 Manhattan Distance is calculated

*/

func ManhattanDistance(puzzle, goal, strNumber string) int {

	i, j := getCoordinates(puzzle, strNumber)
	x, y := getCoordinates(goal, strNumber)

	if !(i == x && j == y) {

		return int(math.Abs(float64(i-x))) + int(math.Abs(float64(j-y)))

	}

	return 0
}

/*

	return moves from source node to goal

	e.g  =>    source node  : 123456078
			   goal  node   : 123456780

	return value  :  right -> right



*/

func Path(solution *Node) string {

	ph := []string{}

	for solution.Parent != nil {

		ph = append(ph, solution.Dir.Lable.String())
		solution = solution.Parent

	}

	return reverse(ph)
}

func reverse(slc []string) string {

	rev_slc := []string{}

	for i := range slc {
		// reverse the order

		rev_slc = append(rev_slc, slc[len(slc)-1-i])

	}

	result := strings.Join(rev_slc, " -> ")

	return result
}

/*

	return value  =>   eg   123456708 -> 123456780

	this function returns all the nodes of the path in the ebove format


*/

func NodesOfPath(solution *Node) string {
	ph := []string{solution.Data}

	for solution.Parent != nil {

		ph = append(ph, solution.Parent.Data)
		solution = solution.Parent

	}

	return reverse(ph)
}

func RemoveIndex(queue []*EghtPuzzle, index int) []*EghtPuzzle {
	return append(queue[:index], queue[index+1:]...)
}
