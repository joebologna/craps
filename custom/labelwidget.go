package custom

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type WidgetTheme struct {
	LabelBorderColor, LabelTextColor color.Color
}

type LabelWidget struct {
	widget.DisableableWidget
	label       *canvas.Text
	border      *canvas.Rectangle
	widgetTheme WidgetTheme
	inverted    bool
}

func NewLabelWidget(text string, widgetTheme WidgetTheme, inverted bool) *LabelWidget {
	var label = canvas.NewText(text, color.White)
	var border = canvas.NewRectangle(color.Transparent)
	setTheme(label, border, widgetTheme, inverted)
	w := &LabelWidget{label: label, border: border, widgetTheme: widgetTheme, inverted: inverted}
	w.ExtendBaseWidget(w)
	return w
}

func (w *LabelWidget) Refresh() {
	setTheme(w.label, w.border, w.widgetTheme, w.inverted)
	w.label.Refresh()
	w.border.Refresh()
	w.BaseWidget.Refresh()
}

func setTheme(label *canvas.Text, border *canvas.Rectangle, widgetTheme WidgetTheme, inverted bool) {
	border.FillColor = color.Transparent
	border.StrokeWidth = 2
	label.Alignment = fyne.TextAlignCenter
	isDark := fyne.CurrentApp().Settings().ThemeVariant() == theme.VariantDark
	if isDark {
		if inverted {
			border.FillColor = widgetTheme.LabelBorderColor
			border.StrokeColor = widgetTheme.LabelTextColor
			label.Color = color.Black
		} else {
			label.Color = widgetTheme.LabelBorderColor
			border.StrokeColor = widgetTheme.LabelTextColor
		}
	} else {
		if inverted {
			border.FillColor = widgetTheme.LabelBorderColor
			border.StrokeColor = color.Black
			label.Color = color.White
		} else {
			label.Color = widgetTheme.LabelBorderColor
			border.StrokeColor = color.Black
		}
	}
	border.SetMinSize(fyne.NewSize(float32(len(label.Text)*int(label.TextSize/2)), float32(label.TextSize)*2))
}

func (w *LabelWidget) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewStack(w.border, container.NewPadded(w.label))
	return widget.NewSimpleRenderer(c)
}
