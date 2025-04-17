package custom

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ButtonWidget struct {
	widget.DisableableWidget
	button      *widget.Button
	border      *canvas.Rectangle
	widgetTheme WidgetTheme
	inverted    bool
}

func NewButtonWidget(text string, widgetTheme WidgetTheme, inverted bool, tapped func()) *ButtonWidget {
	var button = widget.NewButton(text, tapped)
	var border = canvas.NewRectangle(color.Transparent)
	setButtonTheme(button, border, widgetTheme, inverted)
	w := &ButtonWidget{button: button, border: border}
	w.ExtendBaseWidget(w)
	return w
}

func (w *ButtonWidget) Refresh() {
	setButtonTheme(w.button, w.border, w.widgetTheme, w.inverted)
	w.button.Refresh()
	w.border.Refresh()
	w.BaseWidget.Refresh()
}

func setButtonTheme(_ *widget.Button, border *canvas.Rectangle, widgetTheme WidgetTheme, inverted bool) {
	border.FillColor = color.Transparent
	border.StrokeWidth = 2
	isDark := fyne.CurrentApp().Settings().ThemeVariant() == theme.VariantDark
	if isDark {
		if inverted {
			border.FillColor = widgetTheme.LabelBorderColor
			border.StrokeColor = widgetTheme.LabelTextColor
		} else {
			border.StrokeColor = widgetTheme.LabelTextColor
		}
	} else {
		if inverted {
			border.FillColor = widgetTheme.LabelBorderColor
			border.StrokeColor = color.Black
		} else {
			border.StrokeColor = color.Black
		}
	}
}

func (w *ButtonWidget) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewStack(w.border, container.NewPadded(w.button))
	return widget.NewSimpleRenderer(c)
}
