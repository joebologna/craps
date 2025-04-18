package apps

import (
	"craps/custom"
	"embed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func App3(animationFiles embed.FS) *fyne.Container {
	theme1 := custom.WidgetTheme{LabelBorderColor: custom.GREEN, LabelTextColor: custom.OFF_WHITE}
	btn := custom.NewButtonWidget("Green Border", theme1, func() {})
	return container.NewCenter(btn)
}
