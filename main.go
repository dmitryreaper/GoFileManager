package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := createMainWindow(myApp)
	myWindow.ShowAndRun()
}
