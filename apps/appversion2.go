package apps

import (
	"craps/custom"
	"embed"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func App2(animationFiles embed.FS) *fyne.Container {
	var label *custom.CustomLabel
	label = custom.NewCustomLabel("Dark Pref Label", func() {
		if fyne.CurrentApp().Settings().ThemeVariant() == theme.VariantDark {
			label.Text.Color = color.RGBA{255, 0, 0, 255}
			label.Text.Text = "Dark Pref Label = Dark Mode"
		} else {
			label.Text.Color = color.RGBA{255, 0, 0, 255}
			label.Text.Text = "Dark Pref Label = Light Mode"
		}
	})

	var button = custom.NewCustomButton("Button", func() {}, nil)
	button.Tracker = func() {
		if fyne.CurrentApp().Settings().ThemeVariant() == theme.VariantDark {
			button.Overlay.StrokeColor = color.White
		} else {
			button.Overlay.StrokeColor = color.Black
		}
	}
	return container.NewCenter(label.Text)
}
