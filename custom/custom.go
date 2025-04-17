package custom

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type CustomLabel struct {
	*canvas.Text
	Tracker func()
}

func NewCustomLabel(text string, tracker func()) *CustomLabel {
	l := &CustomLabel{canvas.NewText(text, color.White), tracker}
	l.Alignment = fyne.TextAlignCenter

	fyne.CurrentApp().Lifecycle().SetOnStarted(func() {
		l.Tracker()
	})

	fyne.CurrentApp().Lifecycle().SetOnEnteredForeground(func() {
		l.Tracker()
	})

	return l
}

type CustomButton struct {
	*widget.Button
	Overlay *canvas.Rectangle
	Tracker func()
	fyne.CanvasObject
}

func NewCustomButton(text string, tapped, tracker func()) *CustomButton {
	o := canvas.NewRectangle(color.Transparent)
	o.StrokeWidth = 2
	o.StrokeColor = color.White
	o.CornerRadius = 8
	var c fyne.CanvasObject
	b := &CustomButton{widget.NewButton(text, tapped), o, tracker, c}
	b.CanvasObject = container.NewStack(b.Button, b.Overlay)

	fyne.CurrentApp().Lifecycle().SetOnStarted(func() {
		b.Tracker()
	})

	fyne.CurrentApp().Lifecycle().SetOnEnteredForeground(func() {
		b.Tracker()
	})

	return b
}
