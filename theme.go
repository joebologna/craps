package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// CustomTheme defines a custom theme with a larger font size
type CustomTheme struct {
	fyne.Theme
}

func (c CustomTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 18
	}
	return c.Theme.Size(name)
}

func (c CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return c.Theme.Color(name, variant)
}
