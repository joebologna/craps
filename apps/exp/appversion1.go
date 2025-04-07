package exp

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Generates a grid with entries and a button which updates them.
//
// Issues: The AdaptiveGrid does not behave as expected when rotating the phone. It does change the layout of the grid,
// but the device orientation is always vertical on mobile and horizontal left on desktop (regardless of screenSize)
func App1() (*fyne.Container, *widget.Button) {
	button := widget.NewButton("update rows", func() { os.Exit(0) })
	button.Alignment = widget.ButtonAlignLeading
	button.Importance = widget.HighImportance

	stuff := container.NewVBox(widget.NewLabel("hi"))
	return stuff, button
}
