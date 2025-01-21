package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("File Manager")

	currentDir := "."

    // list files and directory
    fileList := widget.NewList(
        func() int {
            files, _ := ioutil.ReadDir(".")
            return len(files)
        },
        func() fyne.CanvasObject {
            return widget.NewLabel("")
        },
        func(i int, o fyne.CanvasObject) {
            files, _ := ioutil.ReadDir(".")
            o.(*widget.Label).SetText(files[i].Name())
        })

	dirLabel := widget.NewLabel(fmt.Sprintf("Directory: %s", currentDir))

	// push buttom
    fileList.OnSelected = func(id int) {
        files, _ := ioutil.ReadDir(".")
        selectedFile := files[id]
        if selectedFile.IsDir() {
            os.Chdir(selectedFile.Name())
            fileList.Refresh()
        }
    }

    // back
    backButton := widget.NewButton("Back", func() {
        os.Chdir("..")
        fileList.Refresh()
    })

    // container
    content := container.NewBorder(
		container.NewVBox(fileList, backButton), 
		dirLabel,
		nil,
		nil,
		fileList,
	)
	fileList.Resize(fyne.NewSize(400, 300))
    myWindow.SetContent(content)

    myWindow.Resize(fyne.NewSize(400, 300))
    myWindow.ShowAndRun()
}
