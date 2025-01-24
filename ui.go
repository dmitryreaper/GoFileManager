package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
)

var Content fyne.CanvasObject

func createMainWindow(myApp fyne.App) fyne.Window {
	myWindow := myApp.NewWindow("File Manager")

	currentDir, _ := os.Getwd()
	dirLabel := widget.NewLabel(fmt.Sprintf("Directory: %s", currentDir))
	fileList, backButton := createFileList(currentDir, dirLabel, myWindow)

	Content = container.NewBorder(
		container.NewVBox(fileList, backButton),
		dirLabel,
		nil,
		nil,
		fileList,
	)

	myWindow.SetContent(Content)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.SetFixedSize(true)
	myWindow.CenterOnScreen()
	return myWindow
}
