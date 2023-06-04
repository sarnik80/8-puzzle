package eightpuzzle

//reachable state for 8puzzle is  =>   181,440

/*

	Each puzzle is a node


*/

type Node struct {
	Data    string
	Level   int32
	Dir     Direction
	G_score int32
	H_score int32
	F_score int32
	Parent  *Node
}

func CreateNode(data string, level int32, g_score, h_score, f_score int32, parent *Node) *Node {

	newNode := Node{}

	newNode.Data = data
	newNode.Level = level
	newNode.G_score = g_score
	newNode.H_score = h_score
	newNode.Parent = parent

	return &newNode

}
