package algorithm

import (
	eightpuzzle "github.com/sarnik80/8-puzzle/eightPuzzle"
)

/*




 */

type BFS struct {
	Name algorithmName
}

func GetName() algorithmName {

	return Bfs
}

func (b BFS) Solve(sourcePuzzle, goalPuzzle string) *eightpuzzle.EghtPuzzle {

	node := &eightpuzzle.Node{Data: sourcePuzzle, Parent: nil, Level: 0}

	source := eightpuzzle.EghtPuzzle{State: node, GoalState: goalPuzzle}

	queue := []*eightpuzzle.EghtPuzzle{&source}
	visited := map[string]eightpuzzle.EghtPuzzle{sourcePuzzle: source}

	for len(queue) > 0 {

		currentPuzzle := queue[0]
		queue = eightpuzzle.RemoveIndex(queue, 0)

		if currentPuzzle.IsGoal() {

			solution := currentPuzzle

			return solution
		}

		visited[currentPuzzle.State.Data] = *currentPuzzle

		for _, child := range *currentPuzzle.GetChildren() {

			_, check := visited[child.State.Data]

			if check == false {

				queue = append(queue, child)

			}

		}

	}

	return nil

}
