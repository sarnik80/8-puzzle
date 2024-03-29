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


	** space complexity  of bfs  O(b^(d+1))   => b is branching factor  and d is depth

	**if there is a solution , bfs will definitely find it
	=> is a complete algorithm


	** if there is more than one solution bfs can find the
	minimal one that  requires less  number of steps  => optimal algorithm



	The number of nodes genreted by this algorithm   :=>     b  +  b ^ 2  +  ....  b ^ d  + b ^(d+1)  - b







*/

type BFS struct {
	Name algorithmName
}

func (b BFS) GetName() algorithmName {

	return Bfs
}

func (b BFS) Solve(sourcePuzzle, goalPuzzle string) (puzzle *eightpuzzle.EghtPuzzle, pop_nodes int, visitedNodes int) {

	/*

			 create source node and source puzzle
		   	 start node     data = sourcpuzzle   ,  level = 0 , g_score = 0 ,  h_score = 0 , f_score = 0  parent = nil

	*/
	node := eightpuzzle.CreateNode(sourcePuzzle, 0, 0, 0, 0, nil)
	source := eightpuzzle.EghtPuzzle{State: node, GoalState: goalPuzzle} // source puzzle

	//  queue of  nodes and map of visited nodes
	queue := []*eightpuzzle.EghtPuzzle{&source}
	visited := make(map[string]bool)

	pop_nodes = 0
	for len(queue) > 0 {

		//  Among all the nodes in the queue, we delete the first node.
		currentPuzzle := queue[0]
		queue = queue[1:] // remove last node
		pop_nodes++

		// End of search
		if currentPuzzle.IsGoal() {

			return currentPuzzle, pop_nodes, len(visited) // target node
		}

		// The current node is added to the visited nodes
		visited[currentPuzzle.State.Data] = true

		// Extending  the current node
		for _, child := range currentPuzzle.GetChildren() {

			check := visited[child.State.Data]

			if !check {

				queue = append(queue,
					child)

			}

		}

	}

	return nil, pop_nodes, len(visited)

}
