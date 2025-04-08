package exp

import (
	"bytes"
	"craps/opts"
	"embed"
	"fmt"
	"image"
	"image/png"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func App2(animationFiles embed.FS, opt opts.Options) *fyne.Container {
	images := make([][]image.Image, 6)
	for j := 1; j <= 6; j++ {
		images[j-1] = make([]image.Image, 0)
		for i := 60; i <= 150; i++ {
			fileName := fmt.Sprintf("media/Animation/roll-%d/%04d.png", j, i)
			data, err := animationFiles.ReadFile(fileName)
			if err != nil {
				panic(err)
			}
			img, err := png.Decode(bytes.NewReader(data))
			if err != nil {
				panic(err)
			}
			images[j-1] = append(images[j-1], img)
		}
	}

	var resultString = binding.NewString()
	result := widget.NewLabelWithData(resultString)
	resultString.Set("Please roll.")
	var rollButton *widget.Button
	img := make([]*canvas.Image, 2)
	for i := 0; i < 2; i++ {
		img[i] = canvas.NewImageFromImage(images[i][len(images)-1])
		img[i].FillMode = canvas.ImageFillOriginal
		img[i].ScaleMode = canvas.ImageScaleFastest
		img[i].Refresh()
		// img[i].File = fmt.Sprintf("File[%d][%d]", i, len(images)-1)
	}
	left := 0
	right := 1
	doAnimation := func(tick float32) {
		// there are len(images) to display in 4s, tick will be 0.5 at 2s for instance, which is len(images)/2, so the image # is tick*len(images)
		i := int(tick * float32(len(images[0])-1))
		leftDie := 0
		rightDie := 1
		img[left].Image = images[leftDie][i]
		img[left].FillMode = canvas.ImageFillOriginal
		img[left].ScaleMode = canvas.ImageScaleFastest
		img[left].Refresh()
		img[right].Image = images[rightDie][i]
		img[right].FillMode = canvas.ImageFillOriginal
		img[right].ScaleMode = canvas.ImageScaleFastest
		img[right].Refresh()
		// fmt.Println(i, img[left].File, img[right].File)
		if tick == 1.0 {
			resultString.Set("Please roll.")
		}
	}
	rollButton = widget.NewButton("Roll", func() {
		resultString.Set("Rolling...")
		fyne.NewAnimation(4*time.Second, doAnimation).Start()
	})
	return container.NewVBox(
		container.NewHBox(
			img[left],
			img[right],
		),
		rollButton,
		result,
	)
}
