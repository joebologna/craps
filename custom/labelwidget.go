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
	w := &LabelWidget{label: label, border: border, widgetTheme: widgetTheme, inverted: inverted}
	w.ExtendBaseWidget(w)
	w.setTheme()
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
	w.setTheme()
	w.BaseWidget.Refresh()
}

func (w *LabelWidget) setTheme() {
	if w.widgetTheme.Scale == 0 {
		w.widgetTheme.Scale = 1
	}
	w.border.FillColor = color.Transparent
	w.border.StrokeWidth = 2
	w.label.Alignment = fyne.TextAlignCenter
	// this happens on mobile for some reason
	if w.label.TextSize < 10 {
		w.label.TextSize = 18
	}
	w.label.TextSize = w.label.TextSize * w.widgetTheme.Scale
	isDark := fyne.CurrentApp().Settings().ThemeVariant() == theme.VariantDark
	if isDark {
		if w.inverted {
			w.border.FillColor = w.widgetTheme.LabelTextColor
			w.border.StrokeColor = w.widgetTheme.LabelBorderColor
			w.label.Color = color.Black
		} else {
			w.label.Color = w.widgetTheme.LabelTextColor
			w.border.StrokeColor = w.widgetTheme.LabelBorderColor
		}
	} else {
		if w.inverted {
			w.border.FillColor = w.widgetTheme.LabelTextColor
			w.border.StrokeColor = color.Black
			w.label.Color = color.White
		} else {
			w.label.Color = w.widgetTheme.LabelTextColor
			w.border.StrokeColor = color.Black
		}
	}
	w.border.SetMinSize(fyne.NewSize(float32(len(w.label.Text)*int(w.label.TextSize/2))*w.widgetTheme.Scale, float32(w.label.TextSize)*2*w.widgetTheme.Scale))
	w.border.Refresh()
	w.label.Refresh()
}

func (w *LabelWidget) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewStack(w.border, container.NewPadded(w.label))
	return widget.NewSimpleRenderer(c)
}
