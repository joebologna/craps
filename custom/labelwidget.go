package custom

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type LabelWidget struct {
	widget.DisableableWidget
	label  *canvas.Text
	border *canvas.Rectangle
}

// CreateRenderer implements fyne.Widget.
func (w *LabelWidget) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewStack(container.NewPadded(w.label), w.border)
	return widget.NewSimpleRenderer(c)
}

// Refresh implements fyne.Widget.
// Subtle: this method shadows the method (DisableableWidget).Refresh of LabelWidget.DisableableWidget.
func (w *LabelWidget) Refresh() {
	// th := w.Theme()
	// v := fyne.CurrentApp().Settings().ThemeVariant()
	w.label.Color = theme.Color(theme.ColorNameForeground)
	w.border.StrokeColor, w.border.StrokeWidth = theme.Color(theme.ColorNameForeground), 2
	w.label.Refresh()
	w.border.Refresh()
	w.BaseWidget.Refresh()
}

func NewLabelWidget(text string) *LabelWidget {
	l := canvas.NewText(text, theme.Color(theme.ColorNameForeground))
	b := canvas.NewRectangle(color.Transparent)
	b.StrokeColor, b.StrokeWidth = theme.Color(theme.ColorNameForeground), 2
	w := &LabelWidget{label: l, border: b}
	w.ExtendBaseWidget(w)
	return w
}
