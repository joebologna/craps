package exp

import (
	"craps/opts"
	"embed"
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func App2(animationFiles embed.FS, opt opts.Options) *fyne.Container {
	images := make([]image.Image, 0)
	for i := 60; i <= 150; i++ {
		fileName := fmt.Sprintf("media/Animation/%04d.png", i)
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		img, err := png.Decode(file)
		if err != nil {
			panic(err)
		}
		images = append(images, img)
	}

	var resultString = binding.NewString()
	result := widget.NewLabelWithData(resultString)
	resultString.Set("Please roll.")
	var rollButton *widget.Button
	img := canvas.NewImageFromImage(images[90])
	img.FillMode = canvas.ImageFillOriginal
	img.ScaleMode = canvas.ImageScaleFastest
	img.Refresh()
	doAnimation := func(tick float32) {
		// there are len(images) to display in 4s, tick will be 0.5 at 2s for instance, which is len(images)/2, so the image # is tick*len(images)
		i := int(tick * float32(len(images)-1))
		img.Image = images[i]
		img.FillMode = canvas.ImageFillOriginal
		img.ScaleMode = canvas.ImageScaleFastest
		img.Refresh()
		if tick == 1.0 {
			resultString.Set("Please roll.")
		}
	}
	rollButton = widget.NewButton("Roll", func() {
		resultString.Set("Rolling...")
		fyne.NewAnimation(4*time.Second, doAnimation).Start()
	})
	return container.NewVBox(
		img,
		rollButton,
		result,
	)
}
