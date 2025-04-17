package apps

import (
	"craps/custom"
	"embed"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func App2(a fyne.App, animationFiles embed.FS) *fyne.Container {
	stuff := custom.NewCustomLabel("Dark Pref Label", color.RGBA{255, 0, 0, 255})
	stuff.Tracker = func() {
		if a.Settings().ThemeVariant() == theme.VariantDark {
			stuff.Color = color.White
			stuff.Text.Text = "Dark Pref Label = Dark Mode"
		} else {
			stuff.Color = color.Black
			stuff.Text.Text = "Dark Pref Label = Light Mode"
		}
	}

	a.Lifecycle().SetOnStarted(func() {
		stuff.Tracker()
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		stuff.Tracker()
	})

	return container.NewBorder(nil, nil, nil, nil, stuff.Text)
}
