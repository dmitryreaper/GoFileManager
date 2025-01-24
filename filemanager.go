package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"os"
)

func createFileList(currentDir string, dirLabel *widget.Label,
	mainWindow fyne.Window) (*widget.List, *widget.Button) {
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
		err := os.Chdir(newDir)
		if err != nil {
			return
		}
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
			openTextEditor(currentDir, selectedFile.Name(), currentDir, mainWindow, dirLabel)
		}
	}

	backButton := widget.NewButton("Back", func() {
		updateDir("..")
	})

	return fileList, backButton
}

func createMainContent(currentDir string, mainWindow fyne.Window) fyne.CanvasObject {
	dirLabel := widget.NewLabel(fmt.Sprintf("Directory: %s", currentDir))
	fileList, backButton := createFileList(
		currentDir,
		dirLabel,
		mainWindow)

	mainWindow.SetContent(container.NewVBox(dirLabel,
		container.NewVBox(fileList, backButton)))

	return nil
}

func openTextEditor(currentDir, filename, dir string,
	mainWindow fyne.Window, dirLabel *widget.Label) {

	filePath := dir + string(os.PathSeparator) + filename

	textEditor, err := ioutil.ReadFile(filePath) // textEditor окно где содержимое файла редактируется
	if err != nil {
		dialog.ShowError(err, mainWindow)
		return
	}

	entry := widget.NewMultiLineEntry()
	entry.SetText(string(textEditor))

	saveButton := widget.NewButton("Save", func() {
		err := ioutil.WriteFile(filePath, []byte(entry.Text), 0644)
		if err != nil {
			dialog.ShowError(err, mainWindow)
			return
		}

		dialog.ShowInformation("Done", "File saved", mainWindow)
	})

	backButton := widget.NewButton("Back", func() {
		createMainContent(currentDir, mainWindow)
		mainWindow.SetContent(Content)
	})

	mainWindow.SetContent(
		container.NewBorder(
			dirLabel,
			container.NewVBox(backButton, saveButton),
			nil,
			nil,
			entry,
		),
	)
}
