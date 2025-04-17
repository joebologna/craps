package custom

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

type CustomLabel struct{ *canvas.Text }

func NewCustomLabel(text string) *CustomLabel {
	// assume Dark mode
	var initialMode, curMode = fyne.CurrentApp().Settings().ThemeVariant(), fyne.CurrentApp().Settings().ThemeVariant()

	// Don't know what the color is yet
	label := &CustomLabel{canvas.NewText(text, color.Transparent)}
	label.SetLabelTheme(initialMode)

	// Color matches initial theme
	fyne.CurrentApp().Lifecycle().SetOnStarted(func() {
		if fyne.CurrentApp().Settings().ThemeVariant() != initialMode {
			label.SetLabelTheme(fyne.CurrentApp().Settings().ThemeVariant())
			curMode = fyne.CurrentApp().Settings().ThemeVariant()
		}
	})

	// Color changes when theme has changed (the app must enter background then re-enter foreground for this to happen)
	fyne.CurrentApp().Lifecycle().SetOnEnteredForeground(func() {
		if fyne.CurrentApp().Settings().ThemeVariant() != curMode {
			label.SetLabelTheme(fyne.CurrentApp().Settings().ThemeVariant())
			curMode = fyne.CurrentApp().Settings().ThemeVariant()
		}
	})
	return label
}

func (label *CustomLabel) SetLabelTheme(variant fyne.ThemeVariant) {
	if variant == theme.VariantDark {
		label.Text.Color = color.White // white text on dark background
	} else {
		label.Text.Color = color.Black // black text on light background
	}
}
