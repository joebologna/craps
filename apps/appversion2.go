package apps

import (
	"embed"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func App2(animationFiles embed.FS) *fyne.Container {
	// assume Dark mode
	var initialMode, curMode = fyne.CurrentApp().Settings().ThemeVariant(), fyne.CurrentApp().Settings().ThemeVariant()

	label := canvas.NewText("", color.Transparent)
	setLabelTheme(label, initialMode)
	fyne.CurrentApp().Lifecycle().SetOnStarted(func() {
		if fyne.CurrentApp().Settings().ThemeVariant() != initialMode {
			setLabelTheme(label, fyne.CurrentApp().Settings().ThemeVariant())
			curMode = fyne.CurrentApp().Settings().ThemeVariant()
		}
	})
	fyne.CurrentApp().Lifecycle().SetOnEnteredForeground(func() {
		if fyne.CurrentApp().Settings().ThemeVariant() != curMode {
			setLabelTheme(label, fyne.CurrentApp().Settings().ThemeVariant())
			curMode = fyne.CurrentApp().Settings().ThemeVariant()
		}
	})
	return container.NewCenter(label)
}

func setLabelTheme(label *canvas.Text, variant fyne.ThemeVariant) {
	if variant == theme.VariantDark {
		label.Text = "Dark Mode"
		label.Color = color.White // white text on dark background
	} else {
		label.Text = "Light Mode"
		label.Color = color.Black // black text on light background
	}
}
