package algorithm

import (
	"fmt"
	"sort"

	eightpuzzle "github.com/sarnik80/8-puzzle/eightPuzzle"
	"github.com/sarnik80/8-puzzle/heuristic"
)

type ASTAR struct {
	Name       algorithmName
	InitPuzzle eightpuzzle.EghtPuzzle
	Heuristic  heuristic.Heuristic
}

func (A ASTAR) GetName() algorithmName {

	return AStar
}

func (aStar ASTAR) solve(sourceState, goalState string) *eightpuzzle.EghtPuzzle {

	i, j := eightpuzzle.FindIndex(sourceState, eightpuzzle.ZeroStr)

	cost := i*3 + j

	h_value := aStar.Heuristic.H_value(sourceState, goalState)

	n := eightpuzzle.Node{Data: sourceState, G_score: cost, F_score: h_value + cost, Parent: nil}

	source := eightpuzzle.EghtPuzzle{State: &n, GoalState: goalState}

	queue := []*eightpuzzle.EghtPuzzle{&source}
	visited := map[string]bool{}

	pop_num := 0

	for len(queue) != 0 {

		currentPuzzle := queue[0]
		queue = eightpuzzle.RemoveIndex(queue, 0)
		pop_num += 1

		if currentPuzzle.IsGoal() {

			fmt.Println("OOOOO", len(queue))
			fmt.Print(">>>visited : ", len(visited))
			return currentPuzzle

		}

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

	return nil
}
