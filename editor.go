package main

import (
    "io/ioutil"
    "os"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/widget"
)

// openTextEditor открывает текстовый редактор для выбранного файла
func openTextEditor(currentDir, filename, dir string, mainWindow fyne.Window, dirLabel *widget.Label) {
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
        mainWindow.SetContent(createMainContent(currentDir, mainWindow)) // исправлено
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
