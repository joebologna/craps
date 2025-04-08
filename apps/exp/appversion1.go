package exp

import (
	"bytes"
	"embed"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// stub
func App1(animationFiles embed.FS) *fyne.Container {
	images := make([]fyne.CanvasObject, 0)
	for i := range []int{60, 150} {
		fileName := fmt.Sprintf("media/Animation/%04d.png", i)
		data, err := animationFiles.ReadFile(fileName)
		if err == nil {
			img := canvas.NewImageFromReader(bytes.NewReader(data), fileName)
			img.FillMode = canvas.ImageFillOriginal
			img.Hide()
			images = append(images, img)
		}
	}
	return container.NewCenter(images...)
}
