package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func createFileList(currentDir string, dirLabel *widget.Label) (*widget.List, *widget.Button) {
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
		}
	}

	backButton := widget.NewButton("Back", func() {
		updateDir("..")
	})

	return fileList, backButton
}
