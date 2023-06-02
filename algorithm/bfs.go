package algorithm

import (
	eightpuzzle "github.com/sarnik80/8-puzzle/eightPuzzle"
)

/*

	[bfs] is uninformed search algorithm that used
	to search a tree or graph data structure .

	it starts at thetree's root of graph and
	visits all nondes at the current depth
	level before moving on the nodes at
	next depth level


	to avoid processing one node more than once
	we divide the nodes into two type :

	1. visited
	2. not visited

	1. visited is implemented by map[current_node]boolean  data type

	2. BFS uses queue data structure for traversal


	** time complextiry  = >   O(V + E)  =>

		V is number of nodes
		E is number of edges


	** space complexity  of bfs  O(b^d)   => b is branching factor  and d is depth

	**if there is a solution , bfs will definitely find it
	=> is a complete algorithm


	** if there is more than one solution bfs can find the
	minimal one that  requires less  number of steps  => optimal algorithm





*/

type BFS struct {
	Name algorithmName
}

func GetName() algorithmName {

	return Bfs
}

func (b BFS) Solve(sourcePuzzle, goalPuzzle string, ch chan string) {

	// create source node and source puzzle

	node := &eightpuzzle.Node{Data: sourcePuzzle, Parent: nil, Level: 0}

	source := eightpuzzle.EghtPuzzle{State: node, GoalState: goalPuzzle}

	queue := []*eightpuzzle.EghtPuzzle{&source}
	visited := map[string]bool{sourcePuzzle: true}

	for len(queue) > 0 {

		//  Among all the nodes in the queue, we delete the first node.

		currentPuzzle := queue[0]
		queue = eightpuzzle.RemoveIndex(queue, 0)

		// End of search

		if currentPuzzle.IsGoal() {

			ch <- eightpuzzle.Path(currentPuzzle.State) // target node
		}

		// The current node is added to the visited nodes

		visited[currentPuzzle.State.Data] = true

		// Extending  the current node

		for _, child := range *currentPuzzle.GetChildren() {

			check := visited[child.State.Data]

			if check == false {

				queue = append(queue, child)

			}

		}

	}

	ch <- "No"

}
