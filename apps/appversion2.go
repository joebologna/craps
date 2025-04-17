package apps

import (
	"craps/custom"
	"embed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func App2(animationFiles embed.FS) *fyne.Container {
	return container.NewCenter(custom.NewLabelWidget("This is a label widget"))
}
