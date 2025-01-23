package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func createFileList(currentDir string, dirLabel *widget.Label, mainWindow fyne.Window) (*widget.List, *widget.Button) {
	fileList := widget.NewList(
		func() int {
			files, _ := ioutil.ReadDir(currentDir)
			return len(files)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i int, o fyne.CanvasObject) {
			
			files, _ := ioutil.ReadDir(currentDir)
			o.(*widget.Label).SetText(files[i].Name())
		},
	)

	updateDir := func(newDir string) {
		os.Chdir(newDir)
		currentDir, _ = os.Getwd()
		dirLabel.SetText(fmt.Sprintf("Directory: %s", currentDir))
		fileList.Refresh()
	}

	fileList.OnSelected = func(id int) {
		files, _ := ioutil.ReadDir(currentDir)
		selectedFile := files[id]
		if selectedFile.IsDir() {
			updateDir(selectedFile.Name())
		} else {
			OpenTextEditor(selectedFile.Name(), currentDir, mainWindow)
		}
	}

	backButton := widget.NewButton("Back", func() {
		updateDir("..")
	})

	return fileList, backButton
}

func createMainContent(currentDir string, mainWindow fyne.Window) fyne.CanvasObject {
	dirLabel := widget.NewLabel(fmt.Sprintf("Directory: %s", currentDir))
	fileList, backButton := createFileList(currentDir, dirLabel, mainWindow)

	return container.NewVBox(
		dirLabel,
		container.NewHBox(backButton),
		fileList,
	)
}
