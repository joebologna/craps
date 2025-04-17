package custom

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

type CustomLabel struct {
	*canvas.Text
	overlay *canvas.Rectangle
	Stack   fyne.CanvasObject
}

func NewCustomLabel(text string) *CustomLabel {
	// assume Dark mode
	var initialMode, curMode = fyne.CurrentApp().Settings().ThemeVariant(), fyne.CurrentApp().Settings().ThemeVariant()

	// we are just looking to put a border around the button
	overlay := canvas.NewRectangle(color.Transparent)
	// Don't know what the color is yet
	overlay.StrokeColor, overlay.StrokeWidth, overlay.CornerRadius = color.Transparent, 2, 4
	overlay.SetMinSize(fyne.NewSize(400, 100))

	// Don't know what the color is yet
	label := &CustomLabel{canvas.NewText(text, theme.Color(theme.ColorNameForeground)), overlay, nil}
	label.SetLabelTheme()
	label.Stack = container.NewStack(label.Text, label.overlay)

	// Color matches initial theme
	fyne.CurrentApp().Lifecycle().SetOnStarted(func() {
		if fyne.CurrentApp().Settings().ThemeVariant() != initialMode {
			label.SetLabelTheme()
			curMode = fyne.CurrentApp().Settings().ThemeVariant()
		}
	})

	// Color changes when theme has changed (the app must enter background then re-enter foreground for this to happen)
	fyne.CurrentApp().Lifecycle().SetOnEnteredForeground(func() {
		if fyne.CurrentApp().Settings().ThemeVariant() != curMode {
			label.SetLabelTheme()
			curMode = fyne.CurrentApp().Settings().ThemeVariant()
		}
	})
	return label
}

func (label *CustomLabel) SetLabelTheme() {
	label.Alignment = fyne.TextAlignCenter
	if fyne.CurrentApp().Settings().ThemeVariant() == theme.VariantDark {
		label.Text.Color = theme.Color(theme.ColorNameForeground)
		label.overlay.StrokeColor = theme.Color(theme.ColorNameForeground)
	} else {
		label.Text.Color = color.Black
		label.overlay.StrokeColor = color.Black
	}
	label.overlay.FillColor = color.Transparent
}
