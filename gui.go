package main

import (
	"strconv"

	// importing fyne
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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

	window.Resize(fyne.NewSize(1200, 600))

	//set up the main menu
	mainMenu := createMenuItems()
	window.SetMainMenu(mainMenu)

	playGround := createPlayGround()

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

	/*

	 create and use button  to start processing
	 first value is button name
	 second value is an action

	*/

	startBTN := widget.NewButton("Start", func() {})

	// using our widgets on our window  (Setup content)

	window.SetContent(container.NewHSplit(

		playGround, // add grid

		container.NewVBox(darkMod, select_entry, startBTN),
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

func createPlayGround() *fyne.Container {

	playGround := container.NewGridWithColumns(3)

	/*
			use for loop to create playGround
		 	our playground is 3 * 3

	*/

	for i := 0; i < 9; i++ {

		// we use Entry widget for each cell

		entrW := widget.NewEntry()

		// setting placeHolder for Each cell

		entrW.SetPlaceHolder(strconv.Itoa(i))

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

	select_entry := widget.NewSelectEntry([]string{"A*", "IDS", "BFS", "IDA*"})

	// what we want to do with selected entry ?!

	select_entry.OnSubmitted = func(s string) {}

	//set default place holder for select_entry wiidget

	select_entry.SetPlaceHolder("AI Algorithm")

	return select_entry

}