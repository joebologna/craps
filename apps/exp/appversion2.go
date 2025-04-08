package exp

import (
	"bytes"
	"craps/opts"
	"embed"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func App2(animationFiles embed.FS, opt opts.Options) *fyne.Container {
	images := make([]*canvas.Image, 0)
	for i := 60; i <= 150; i++ {
		fileName := fmt.Sprintf("media/Animation/%04d.png", i)
		data, err := animationFiles.ReadFile(fileName)
		if err == nil {
			tmp := canvas.NewImageFromReader(bytes.NewReader(data), fileName)
			tmp.FillMode = canvas.ImageFillOriginal
			images = append(images, tmp)
		} else {
			panic(err)
		}
	}

	var resultString = binding.NewString()
	result := widget.NewLabelWithData(resultString)
	resultString.Set("Please roll.")
	var rollButton *widget.Button
	fileName := fmt.Sprintf("media/Animation/%04d.png", 150)
	data, err := animationFiles.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	img := canvas.NewImageFromReader(bytes.NewReader(data), fileName)
	img.FillMode = canvas.ImageFillOriginal
	img.ScaleMode = canvas.ImageScaleFastest
	doAnimation := func(tick float32) {
		// there are len(images) to display in 4s, tick will be 0.5 at 2s for instance, which is len(images)/2, so the image # is tick*len(images)
		i := int(tick * float32(len(images)-1))
		img.Resource = images[i].Resource
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
