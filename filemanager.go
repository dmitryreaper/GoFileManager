package main

import (
    "fmt"
    "io/ioutil"
    "os"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

// createFileList создает список файлов в текущей директории
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

// createMainContent создает основное содержимое окна с файловым менеджером
func createMainContent(currentDir string, mainWindow fyne.Window) fyne.CanvasObject {
    dirLabel := widget.NewLabel(fmt.Sprintf("Directory: %s", currentDir))
    fileList, backButton := createFileList(currentDir, dirLabel, mainWindow)
    mainWindow.SetContent(container.NewVBox(dirLabel,
        container.NewVBox(fileList, backButton)))
    return nil
}
