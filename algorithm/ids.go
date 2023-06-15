package algorithm

import eightpuzzle "github.com/sarnik80/8-puzzle/eightPuzzle"

type IDS struct {
	Name       algorithmName
	DepthLimit int
}

func (ids IDS) GetName() algorithmName {

	return Ids
}

func (ids IDS) Solve(sourcePuzzle, goalPuzzle string) *eightpuzzle.EghtPuzzle {

	for depth := 0; depth < ids.DepthLimit; depth++ {

		result := dfs(sourcePuzzle, goalPuzzle, ids.DepthLimit)

		if result != nil {

			return result
		}

	}

	return nil
}

func dfs(sourcePuzzle, goalPuzzle string, depthLimit int) *eightpuzzle.EghtPuzzle {

	node := eightpuzzle.CreateNode(sourcePuzzle, 0, 0, 0, 0, nil)
	source := eightpuzzle.EghtPuzzle{State: node, GoalState: goalPuzzle} // source puzzle

	stack := []eightpuzzle.EghtPuzzle{source}
	visited := map[eightpuzzle.Node]bool{}

	for len(stack) > 0 {

		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.IsGoal() {

			return &current
		}

		if !visited[*current.State] {
			visited[*current.State] = true
		}

		for _, child := range current.GetChildren() {

			if child.State.Level < depthLimit {

				if !visited[*child.State] {

					stack = append(stack, *child)

				}
			}
		}

	}

	return nil
}
