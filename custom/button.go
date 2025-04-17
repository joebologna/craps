package custom

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type CustomButton struct {
	*widget.Button
	overlay *canvas.Rectangle
	Stack   fyne.CanvasObject
}

func NewCustomButton(text string, tapped func()) *CustomButton {
	// assume Dark mode
	var initialMode, curMode = fyne.CurrentApp().Settings().ThemeVariant(), fyne.CurrentApp().Settings().ThemeVariant()

	// we are just looking to put a border around the button
	overlay := canvas.NewRectangle(color.Transparent)
	// Don't know what the color is yet
	overlay.StrokeColor, overlay.StrokeWidth, overlay.CornerRadius = color.Transparent, 2, 4
	overlay.SetMinSize(fyne.NewSize(400, 100))

	button := &CustomButton{widget.NewButton(text, tapped), overlay, nil}
	button.SetLabelTheme()
	button.Stack = container.NewStack(button.Button, button.overlay)

	// Color matches initial theme
	fyne.CurrentApp().Lifecycle().SetOnStarted(func() {
		if fyne.CurrentApp().Settings().ThemeVariant() != initialMode {
			button.SetLabelTheme()
			curMode = fyne.CurrentApp().Settings().ThemeVariant()
		}
	})

	// Color changes when theme has changed (the app must enter background then re-enter foreground for this to happen)
	fyne.CurrentApp().Lifecycle().SetOnEnteredForeground(func() {
		if fyne.CurrentApp().Settings().ThemeVariant() != curMode {
			button.SetLabelTheme()
			curMode = fyne.CurrentApp().Settings().ThemeVariant()
		}
	})
	return button
}

// var red = color.RGBA{255, 0, 0, 255}

func (button *CustomButton) SetLabelTheme() {
	button.overlay.FillColor = color.Transparent
	button.overlay.StrokeColor = theme.Color(theme.ColorNameForeground)
}
