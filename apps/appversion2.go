package apps

import (
	"craps/custom"
	"craps/utils"
	"embed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func App2(animationFiles embed.FS) *fyne.Container {
	theme1 := custom.WidgetTheme{LabelBorderColor: custom.GREEN, LabelTextColor: custom.OFF_WHITE}
	theme2 := custom.WidgetTheme{LabelBorderColor: custom.RED, LabelTextColor: custom.OFF_WHITE, Scale: 0.75}
	l1 := custom.NewLabelWidget("This is a label widget", theme1, false)
	l2 := custom.NewLabelWidget("This is another label widget", theme1, true)
	l2t := utils.NewBS()
	l2t.Set("This is a label widget")
	l3 := custom.NewLabelWidgetWithData(l2t, theme2, false)
	l3t := utils.NewBS()
	l3t.Set("This is another label widget")
	l4 := custom.NewLabelWidgetWithData(l3t, theme2, true)
	rc := container.NewCenter(container.NewVBox(
		container.NewGridWithRows(2, l1, l2),
		container.NewGridWithRows(2, l3, l4),
		custom.NewButtonWidget("Push me", theme2, false, func() {
			l3t.Set("pushed")
		}),
	))
	return rc
}
