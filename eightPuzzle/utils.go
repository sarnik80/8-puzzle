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

func findIndex(puzzle, numbrt string) (i, j int) {

	index := strings.Index(puzzle, numbrt)

	i = index / maxRow

	j = index - (i * maxCol)

	return

}

//  Calculates the coordinates of a number in a puzzle

func getCoordinates(puzzle, strNumber string) (x, y int) {

	x, y = findIndex(puzzle, strNumber)

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

func RemoveIndex(queue []*EghtPuzzle, index int) []*EghtPuzzle {
	return append(queue[:index], queue[index+1:]...)
}

/*

 By having the coordinates of a tile in the
 current puzzle and in the target puzzle,
 Manhattan Distance is calculated

*/

func manhattanDistance(puzzle, goal, strNumber string) float64 {

	i, j := getCoordinates(puzzle, strNumber)
	x, y := getCoordinates(goal, strNumber)

	if !(i == x && j == y) {

		return math.Abs(float64(i)-float64(x)) + math.Abs(float64(j)-float64(y))

	}

	return 0
}

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

	result := strings.Join(rev_slc, "->")

	return result
}
