package exp

import (
	"embed"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// stub
func App1(_ embed.FS) (*fyne.Container, *widget.Button) {
	button := widget.NewButton("update rows", func() { os.Exit(0) })
	button.Alignment = widget.ButtonAlignLeading
	button.Importance = widget.HighImportance

	stuff := container.NewVBox(widget.NewLabel("hi"))
	return stuff, button
}
