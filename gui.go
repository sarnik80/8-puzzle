package main

import (
	"fmt"
	"image/color"
	"strings"

	// importing fyne
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/sarnik80/8-puzzle/algorithm"
	eightpuzzle "github.com/sarnik80/8-puzzle/eightPuzzle"
	"github.com/sarnik80/8-puzzle/heuristic"
)

func createAndShowMyApp() {

	// this line will create new App

	myApp := app.New()

	//we will create a new window  and we set a title for our app

	window := myApp.NewWindow("8-Puzzle")

	/*

			Resizing our Fyne app window
		    first one : width
		    2nd one   : height


	*/

	window.Resize(fyne.NewSize(1000, 500))

	//set up the main menu
	mainMenu := createMenuItems()
	window.SetMainMenu(mainMenu)

	playGround := createPlayGround(strings.Split(eightpuzzle.GoalPuzzle, ""))

	select_entry := createSelectEntry()

	/*

	  this is a checkbox widget
	  we can change theme from Dark to Light

	*/

	darkMod := widget.NewCheck("Light", func(b bool) {

		if !b {
			myApp.Settings().SetTheme(theme.DarkTheme())
		} else {
			myApp.Settings().SetTheme(theme.LightTheme())
		}

	})

	v := container.NewVBox()

	entrW := widget.NewEntry()

	entrW.SetPlaceHolder(fmt.Sprintf("e.g %s", eightpuzzle.GoalPuzzle))

	/*

	 create and use button  to start processing
	 first value is button name
	 second value is an action

	*/

	startBTN := widget.NewButton("Start", func() {

		if select_entry.Text == "" || entrW.Text == "" {

			fmt.Println("choose one strategy OR enter your puzzle sequence")
		} else {

			resultPage := myApp.NewWindow(select_entry.Text)

			switch select_entry.Text {

			case algorithm.AStar.String():

				entrW.SetText("")
				entrW.Refresh()

				select_entry.SetText("")
				select_entry.Refresh()

				content := callAStar(entrW.Text, eightpuzzle.GoalPuzzle)

				resultPage.SetContent(content)

				resultPage.Show()

			case algorithm.Bfs.String():

				result := callBFS(entrW.Text, eightpuzzle.GoalPuzzle)

				entrW.SetText("")
				entrW.Refresh()

				select_entry.SetText("")
				select_entry.Refresh()
				resultPage.SetContent(result)

				resultPage.Show()

			}

		}
	})

	v.Objects = append(v.Objects)

	// using our widgets on our window  (Setup content)

	window.SetContent(container.NewHSplit(

		container.NewVBox(playGround, darkMod, select_entry, entrW, startBTN), // add grid

		v,
	))

	// Finally running our app

	window.ShowAndRun()

}

func createMenuItems() *fyne.MainMenu {

	//First menuItem
	menu := fyne.NewMenu("File")

	mainMenu := fyne.NewMainMenu(menu)

	return mainMenu
}

func createPlayGround(plceHolders []string) *fyne.Container {

	playGround := container.NewGridWithColumns(3)

	/*
			use for loop to create playGround
		 	our playground is 3 * 3

	*/

	for i := 0; i < 9; i++ {

		// we use Entry widget for each cell

		entrW := widget.NewEntry()

		// setting placeHolder for Each cell

		entrW.SetPlaceHolder(plceHolders[i])

		entrW.Disable()

		// appending to the slice of cell or our playground

		playGround.Objects = append(playGround.Objects, entrW)

	}

	return playGround

}

/*

  	 create new selectEntry widget
   	 to choose one of the AI algorithm

*/

func createSelectEntry() *widget.SelectEntry {

	//  it  takes slice of options  []string{A*, IDS, DFS, IDA*}

	select_entry := widget.NewSelectEntry([]string{algorithm.AStar.String(), algorithm.IDAStar.String(), algorithm.Bfs.String(), algorithm.Ids.String()})

	// what we want to do with selected entry ?!

	select_entry.OnSubmitted = func(s string) {}

	//set default place holder for select_entry wiidget

	select_entry.SetPlaceHolder("AI Algorithm")

	return select_entry

}

func callAStar(sourcePuzle, goalPuzzle string) *fyne.Container {

	aStr := algorithm.ASTAR{Name: algorithm.AStar, Heuristic: heuristic.Manhattan{Name: heuristic.ManhattanDistance}}

	resultPuzzle, pop_nodes, visitedNodes := aStr.Solve(sourcePuzle, goalPuzzle)

	depth := resultPuzzle.State.Level
	moves := eightpuzzle.Path(resultPuzzle.State)
	resultNodes := eightpuzzle.NodesOfPath(resultPuzzle.State)

	return createResultPage(resultNodes, moves, pop_nodes, visitedNodes, depth)
}

func callBFS(sourcePuzzle, goalPuzzle string) *fyne.Container {

	bfs := algorithm.BFS{Name: algorithm.Bfs}

	resultPuzzle, pop_nodes, visitedNodes := bfs.Solve(sourcePuzzle, goalPuzzle)
	depth := resultPuzzle.State.Level

	moves := eightpuzzle.Path(resultPuzzle.State)
	resultNodes := eightpuzzle.NodesOfPath(resultPuzzle.State)

	return createResultPage(resultNodes, moves, pop_nodes, visitedNodes, depth)

}

func createResultPage(resultNodes, moves string, pop_nodes, visitedNodes, depth int) *fyne.Container {
	initNode := createPlayGround(strings.Split(strings.Split(resultNodes, "->")[0], ""))

	colorX := color.NRGBA{R: 0, G: 255, B: 0, A: 255}
	moveTXT := canvas.NewText(fmt.Sprintf(">>> Moves : [%s]", moves), colorX)
	resultNodesTXT := canvas.NewText(fmt.Sprintf(">>> Path Nodes : [%s]", resultNodes), colorX)
	popTXT := canvas.NewText(fmt.Sprintf(">>> Poped Nodes : %v", pop_nodes), colorX)
	visitedNodesTXT := canvas.NewText(fmt.Sprintf(">>> Visited Nodes : %v", visitedNodes), colorX)

	depthTXT := canvas.NewText(fmt.Sprintf(">>> Depth : (%v)", depth), colorX)

	vBox := container.NewVBox(initNode, resultNodesTXT, moveTXT, popTXT, visitedNodesTXT, depthTXT)

	return vBox
}
