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
	images := make([]fyne.CanvasObject, 0)
	for i := 60; i <= 150; i++ {
		fileName := fmt.Sprintf("media/Animation/%04d.png", i)
		data, err := animationFiles.ReadFile(fileName)
		if err == nil {
			img := canvas.NewImageFromReader(bytes.NewReader(data), fileName)
			img.FillMode = canvas.ImageFillOriginal
			img.Hide()
			images = append(images, img)
		} else {
			panic(err)
		}
	}

	var resultString = binding.NewString()
	result := widget.NewLabelWithData(resultString)
	resultString.Set("Please roll.")
	var inProgress = false
	var rollButton *widget.Button
	if opt == opts.GoFuncOpt {
		rollButton = widget.NewButton("Roll", func() {
			if !inProgress {
				showImages(images, &inProgress, rollButton, &resultString)
			}
		})
		images[len(images)-1].Show()
		return container.NewVBox(
			container.NewCenter(container.NewStack(images...)),
			rollButton,
			result,
		)
	} else if opt == opts.AnimationWithShowHide {
		images[0].Show()
		doAnimation := func(tick float32) {
			// there are len(images) to display in 4s, tick will be 0.5 at 2s for instance, which is len(images)/2, so the image # is tick*len(images)
			i := int(tick * float32(len(images)-1))
			if i > 0 {
				images[i-1].Hide()
			}
			images[i].Show()
			if tick == 1.0 {
				resultString.Set("Please roll.")
			}
		}
		rollButton = widget.NewButton("Roll", func() {
			resultString.Set("Rolling...")
			images[len(images)-1].Hide()
			fyne.NewAnimation(4*time.Second, doAnimation).Start()
		})
		return container.NewVBox(
			container.NewCenter(container.NewStack(images...)),
			rollButton,
			result,
		)
	} else {
		images[0].Show()
		doAnimation := func(tick float32) {
			// there are len(images) to display in 4s, tick will be 0.5 at 2s for instance, which is len(images)/2, so the image # is tick*len(images)
			i := int(tick * float32(len(images)-1))
			if i > 0 {
				images[i-1].Hide()
			}
			images[i].Show()
			if tick == 1.0 {
				resultString.Set("Please roll.")
			}
		}
		rollButton = widget.NewButton("Roll", func() {
			resultString.Set("Rolling...")
			images[len(images)-1].Hide()
			fyne.NewAnimation(4*time.Second, doAnimation).Start()
		})
		return container.NewVBox(
			container.NewCenter(container.NewStack(images...)),
			rollButton,
			result,
		)
	}
}

func showImages(images []fyne.CanvasObject, inProgress *bool, rollButton *widget.Button, resultString *binding.String) {
	go func() {
		(*resultString).Set("Rolling...")
		*inProgress = true
		rollButton.Disable()
		images[len(images)-1].Hide()
		// images[10].Show()
		// time.Sleep(2 * time.Second)
		for i := 10; i < len(images); i++ {
			images[i].Show()
			time.Sleep(time.Millisecond * 50)
			images[i].Hide()
		}
		images[len(images)-1].Show()
		*inProgress = false
		rollButton.Enable()
		(*resultString).Set("You rolled 3 + 3 = 6")
	}()
}
