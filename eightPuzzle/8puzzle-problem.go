package eightpuzzle

import "reflect"

type EghtPuzzle struct {
	State     *Node
	GoalState string
}

func (e EghtPuzzle) IsGoal() bool {

	return e.State.Data == e.GoalState
}

/*

	This function calls the createMove function for
	all the directions that tile number 0 can move
	and returns all moves for current puzzle

*/

func (e EghtPuzzle) moves() *[]Direction {

	source_i, source_j := e.getTileCoordinates(zeroStr)
	availableDir := *e.getAvailableDirections()
	moves := []Direction{}

	for _, dir := range availableDir {

		move := createMove(source_i, source_j, &dir)

		moves = append(moves, *move)

	}

	return &moves

}

/* returns all available directions of current node   [up , down , left , right]

it depends on zero tile coordinate

 e.g  :

   1 2 3
   5 0 8
   6 7 4

we have 4 available directions to move  [Up, Down , Left , Right]

but  in another egample :

  1 3 4
  5 6 8
  7 2 0

we have 2 available directions to move  => [Up , Left]


*/

func (e EghtPuzzle) getAvailableDirections() *[]Direction {

	// zero tile coordinate
	i, j := e.getTileCoordinates(zeroStr)

	avDirections := findAvailableDirections(i, j)

	return avDirections

}

// It gives us the coordinates of each tile

func (e EghtPuzzle) getTileCoordinates(tile string) (i, j int) {

	i, j = getCoordinates(e.State.Data, tile)

	return
}

/*

	e.g  :

	current node :     1 3 4
  				       5 6 8
  					   7 2 0

	we have 2 move for this puzzle   (Based on tile number 0)

	(0).x => 2     (0).y => 2



	1 . UP  (-1, 0)  :  1 3 4
						5 6 0
						7 2 8

		(0).x => 1   (0).y => 2


	=>

	source_i = 2  ,  source_j = 2

	**We can find the final coordinates of the 0 tile with the following operations :

	dist_i = source_i + UP.X   , dist_j = source_j + UP.Y


*/

func createMove(source_i, source_j int, dir *Direction) *Direction {

	move := Direction{}

	move.X = calCulateDistIndex(source_i, dir.X)
	move.Y = calCulateDistIndex(source_j, dir.Y)

	move.Lable = dir.Lable

	return &move

}

func (e EghtPuzzle) getChildrenData() *map[string]Direction {

	childrenData := map[string]Direction{}

	i, j := e.getTileCoordinates(zeroStr) //  2 , 2

	// Swapper() function is used to swaps the elements of slice

	// zero index of 0 tile in sliced data
	zeroIndex := i*maxRow + j // 8

	moves := *e.moves()

	for _, move := range moves {

		puzzle := slicedString(e.State.Data)
		swap := reflect.Swapper(puzzle)
		// destination index for 0 tile in sliced data
		destIndex := move.X*3 + move.Y // 7

		swap(zeroIndex, destIndex)

		strData := stringedSlice(puzzle)

		childrenData[strData] = move

	}

	return &childrenData
}

/*

	Using the data we obtained, we create nodes for different modes that come up and
	finally return them as children of the current node.


	e.g  =>

	current node  :  1 3 4
  				     5 6 8
  					 7 2 0

    moves  :   [Up , Left]


	childrenData[0]  = first move = up  =  "134560728"   =>   1 3 4
															  5 6 0
															  7 2 8


	childrenData[1]  = second move = left =  "134568702" =>   1 3 4
															  5 6 8
															  7 0 2


	Now we have to create a node from each of these data


	e.g  :
		child1 := Node{Data:= childrenData[0]}


*/

func (e EghtPuzzle) GetChildren() *[]*EghtPuzzle {

	childrenData := e.getChildrenData()
	children := []*EghtPuzzle{}

	for data, dir := range *childrenData {

		child := Node{Data: data, Parent: e.State, Level: e.State.Level + 1, Dir: dir}

		childPuzzle := EghtPuzzle{State: &child, GoalState: e.GoalState}
		children = append(children, &childPuzzle)

	}

	return &children
}
