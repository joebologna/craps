package custom

import (
	"craps/utils"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var RED, GREEN, OFF_WHITE, WHITE = color.RGBA{128, 0, 0, 255}, color.RGBA{0, 128, 0, 255}, color.RGBA{192, 192, 192, 255}, color.RGBA{255, 255, 255, 255}

type WidgetTheme struct {
	LabelBorderColor, LabelTextColor color.Color
	Scale                            float32
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

func NewLabelWidgetWithData(text utils.BS, widgetTheme WidgetTheme, inverted bool) *LabelWidget {
	labelWidget := NewLabelWidget(text.GetS(), widgetTheme, inverted)
	text.AddListener(binding.NewDataListener(func() {
		labelWidget.label.Text = text.GetS()
	}))
	return labelWidget
}

func (w *LabelWidget) Refresh() {
	setTheme(w.label, w.border, w.widgetTheme, w.inverted)
	w.label.Refresh()
	w.border.Refresh()
	w.BaseWidget.Refresh()
}

func setTheme(label *canvas.Text, border *canvas.Rectangle, widgetTheme WidgetTheme, inverted bool) {
	if widgetTheme.Scale == 0 {
		widgetTheme.Scale = 1
	}
	border.FillColor = color.Transparent
	border.StrokeWidth = 2
	label.Alignment = fyne.TextAlignCenter
	// this happens on mobile for some reason
	if label.TextSize < 10 {
		label.TextSize = 18
	}
	label.TextSize = label.TextSize * widgetTheme.Scale
	isDark := fyne.CurrentApp().Settings().ThemeVariant() == theme.VariantDark
	if isDark {
		if inverted {
			border.FillColor = widgetTheme.LabelTextColor
			border.StrokeColor = widgetTheme.LabelBorderColor
			label.Color = color.Black
		} else {
			label.Color = widgetTheme.LabelTextColor
			border.StrokeColor = widgetTheme.LabelBorderColor
		}
	} else {
		if inverted {
			border.FillColor = widgetTheme.LabelTextColor
			border.StrokeColor = color.Black
			label.Color = color.White
		} else {
			label.Color = widgetTheme.LabelTextColor
			border.StrokeColor = color.Black
		}
	}
	border.SetMinSize(fyne.NewSize(float32(len(label.Text)*int(label.TextSize/2))*widgetTheme.Scale, float32(label.TextSize)*2*widgetTheme.Scale))
}

func (w *LabelWidget) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewStack(w.border, container.NewPadded(w.label))
	return widget.NewSimpleRenderer(c)
}
