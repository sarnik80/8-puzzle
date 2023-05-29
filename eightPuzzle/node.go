package eightpuzzle

//reachable state for 8puzzle is  =>   181,440

/*

	Each puzzle is a node


*/

type Node struct {
	Data    string
	Level   int32
	G_score float64
	H_score float64
	F_score float64
	Parent  *Node
}
