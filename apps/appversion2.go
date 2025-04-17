package apps

import (
	"craps/custom"
	"embed"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func App2(animationFiles embed.FS) *fyne.Container {
	theme1 := custom.WidgetTheme{LabelBorderColor: color.RGBA{0, 128, 0, 255}, LabelTextColor: color.White}
	theme2 := custom.WidgetTheme{LabelBorderColor: color.RGBA{128, 0, 0, 255}, LabelTextColor: color.White}
	l1 := custom.NewLabelWidget("This is a label widget", theme1, false)
	l2 := custom.NewLabelWidget("This is another label widget", theme1, true)
	l3 := custom.NewLabelWidget("This is a label widget", theme2, false)
	l4 := custom.NewLabelWidget("This is another label widget", theme2, true)
	rc := container.NewCenter(container.NewVBox(
		container.NewGridWithRows(2, l1, l2),
		container.NewGridWithRows(2, l3, l4),
	))
	return rc
}
