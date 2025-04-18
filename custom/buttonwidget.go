package custom

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ButtonWidget struct {
	widget.DisableableWidget
	button      *widget.Button
	border      *canvas.Rectangle
	widgetTheme WidgetTheme
}

func NewButtonWidget(text string, widgetTheme WidgetTheme, tapped func()) *ButtonWidget {
	var button = widget.NewButton(text, tapped)
	var border = canvas.NewRectangle(color.Transparent)
	setButtonTheme(border, widgetTheme)
	w := &ButtonWidget{button: button, border: border}
	w.ExtendBaseWidget(w)
	return w
}

func (w *ButtonWidget) Enable() {
	w.border.Show()
	w.button.Enable()
}

func (w *ButtonWidget) Disable() {
	w.border.Hide()
	w.button.Disable()
}

func (w *ButtonWidget) Refresh() {
	setButtonTheme(w.border, w.widgetTheme)
	w.button.Refresh()
	w.border.Refresh()
	w.BaseWidget.Refresh()
}

func setButtonTheme(border *canvas.Rectangle, widgetTheme WidgetTheme) {
	border.FillColor = widgetTheme.LabelBorderColor
}

func (w *ButtonWidget) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewStack(w.border, container.NewPadded(w.button))
	return widget.NewSimpleRenderer(c)
}
