package eightpuzzle

//reachable state for 8puzzle is  =>   181,440

/*

	Each puzzle is a node


*/

type Node struct {
	Data    string //123456780
	Level   int
	Dir     Direction
	G_score int
	H_score int
	F_score int
	Parent  *Node
}

//constructor

func CreateNode(data string, level int, g_score, h_score, f_score int, parent *Node) *Node {

	newNode := Node{}

	newNode.Data = data
	newNode.Level = level
	newNode.G_score = g_score
	newNode.H_score = h_score
	newNode.Parent = parent

	return &newNode

}
