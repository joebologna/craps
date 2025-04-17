package apps

import (
	"craps/custom"
	"embed"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func App2(animationFiles embed.FS) *fyne.Container {
	l := custom.NewCustomLabel("Custom Label Widget")

	i1, i2 := 0, 0
	var b1 *custom.CustomButton
	b1 = custom.NewCustomButton("Tracking Theme Variant with Custom Label", func() {
		i1++
		b1.SetText(fmt.Sprintf("tapped %d times", i1))
	})
	var b2 *widget.Button
	b2 = widget.NewButton("just a button", func() {
		i2++
		b2.SetText(fmt.Sprintf("tapped %d times", i2))
	})
	return container.NewCenter(
		container.NewVBox(
			l.Stack,
			custom.NewCustomLabel("Tracking Theme Variant with Custom Label").Text,
			b1.Stack,
			b2,
		),
	)
}
