package algorithm

import (
	"testing"

	eightpuzzle "github.com/sarnik80/8-puzzle/eightPuzzle"
	"github.com/sarnik80/8-puzzle/heuristic"
)

type Tests struct {
	SourcePuzzle string
	resultPuzzle string

	depth int
}

func TestAStar(t *testing.T) {

	tests := []Tests{
		{"123456078", eightpuzzle.GoalPuzzle, 2},
		{"123450678", eightpuzzle.GoalPuzzle, 13},
		{"125340678", eightpuzzle.GoalPuzzle, 21},
		{"125034678", eightpuzzle.GoalPuzzle, 23},
		{"312450678", eightpuzzle.GoalPuzzle, 19},
		{"312045678", eightpuzzle.GoalPuzzle, 21},
	}

	astar := ASTAR{Heuristic: heuristic.Manhattan{Name: heuristic.ManhattanDistance}}

	for _, test := range tests {

		result, _, _ := astar.Solve(test.SourcePuzzle, test.resultPuzzle)

		if result == nil {

			t.Errorf("A*(%v , %v)  FAILED . Expected %v return value is nil!\n\n", test.SourcePuzzle, test.resultPuzzle, test.resultPuzzle)

		} else if result.State.Data != test.resultPuzzle || result.State.Level != test.depth {
			if result.State.Data != test.resultPuzzle {
				t.Errorf("A*(%v , %v)  FAILED . Expected %v return value is %v!\n\n", test.SourcePuzzle, test.resultPuzzle, test.resultPuzzle, result.State.Data)
			}

			if result.State.Level != test.depth {
				t.Errorf("A*(%v , %v)  FAILED . Expected %v return value is %v!\n\n", test.SourcePuzzle, test.resultPuzzle, test.depth, result.State.Level)
			}

		} else {
			t.Logf("A*(%v , %v) PASSED. Expect %v in %v depth  , got  %v  in %v level And Path is : %v\n\n", test.SourcePuzzle, test.resultPuzzle, test.resultPuzzle, test.depth, result.State.Data, result.State.Level, eightpuzzle.Path(result.State))
		}

	}

}
