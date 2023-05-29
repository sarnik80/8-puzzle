package eightpuzzle

type EghtPuzzle struct {
	State     *Node
	GoalState string
}

func (e EghtPuzzle) isGoal() bool {

	return e.State.Data == e.GoalState
}
