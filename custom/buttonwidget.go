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
	w := &ButtonWidget{button: button, border: border, widgetTheme: widgetTheme}
	w.ExtendBaseWidget(w)
	w.setTheme()
	return w
}

func (w *ButtonWidget) Refresh() {
	w.setTheme()
	w.BaseWidget.Refresh()
}

func (w *ButtonWidget) setTheme() {
	w.border.StrokeColor, w.border.StrokeWidth, w.border.CornerRadius = w.widgetTheme.LabelBorderColor, 2, 6
	w.border.Refresh()
}

func (w *ButtonWidget) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewStack(container.NewPadded(w.button), w.border)
	return widget.NewSimpleRenderer(c)
}

func (w *ButtonWidget) Enable() {
	w.border.Show()
	w.button.Enable()
}

func (w *ButtonWidget) Disable() {
	w.border.Hide()
	w.button.Disable()
}
