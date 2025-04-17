package custom

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type CustomLabel struct {
	*canvas.Text
	Tracker func()
}

func NewCustomLabel(text string, c color.Color) *CustomLabel {
	l := &CustomLabel{canvas.NewText(text, c), nil}
	l.Alignment = fyne.TextAlignCenter
	return l
}
