package algorithm

import (
	"sort"

	eightpuzzle "github.com/sarnik80/8-puzzle/eightPuzzle"
	"github.com/sarnik80/8-puzzle/heuristic"
)

type ASTAR struct {
	Name       algorithmName
	InitPuzzle eightpuzzle.EghtPuzzle
	Heuristic  heuristic.Heuristic
}

func (a ASTAR) GetName() algorithmName {

	return AStar
}

func (aStar ASTAR) Solve(sourceState, goalState string) (solution *eightpuzzle.EghtPuzzle, pop_nodes int, visitedNodes int) {

	h_value := aStar.Heuristic.H_value(sourceState, goalState)

	n := eightpuzzle.Node{Data: sourceState, F_score: h_value, Parent: nil}

	source := eightpuzzle.EghtPuzzle{State: &n, GoalState: goalState}

	queue := []*eightpuzzle.EghtPuzzle{&source}
	visited := map[string]bool{}

	pop_nodes = 0

	for len(queue) != 0 {

		currentPuzzle := queue[0]
		pop_nodes++

		if currentPuzzle.IsGoal() {

			return currentPuzzle, pop_nodes, len(visited)

		}

		queue = queue[1:] // remove last node
		visited[currentPuzzle.State.Data] = true

		for _, child := range currentPuzzle.GetChildren() {

			_, check := visited[child.State.Data]

			if !check {

				queue = append(queue, child)

			}

		}

		sort.SliceStable(queue, func(i, j int) bool {

			return aStar.Heuristic.H_value(queue[i].State.Data, queue[i].GoalState)+queue[i].State.G_score < aStar.Heuristic.H_value(queue[j].State.Data, queue[j].GoalState)+queue[j].State.G_score
		})

	}

	return nil, pop_nodes, len(visited)
}
