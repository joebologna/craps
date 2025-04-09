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
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/rand"
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
	rand.Seed(uint64(time.Now().UnixNano()))

	leftDie := 0
	rightDie := 0

	doAnimation := func(tick float32) {
		// there are len(images) to display in 4s, tick will be 0.5 at 2s for instance, which is len(images)/2, so the image # is tick*len(images)
		i := int(tick * float32(len(images[0])-1))
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
			rollButton.Enable()
			total := leftDie + 1 + rightDie + 1 // Dice values are 1-indexed
			resultText := fmt.Sprintf("You rolled: %d", total)
			switch total {
			case 7, 11:
				resultText += ". You Win!"
			case 2, 3, 12:
				resultText += ". You Lose."
			default:
				resultText += ". Roll again."
			}
			resultString.Set(resultText)
		}
	}

	rollButton = widget.NewButton("Roll", func() {
		rollButton.Disable()
		leftDie = rand.Intn(6)
		rightDie = rand.Intn(6)
		resultString.Set("Rolling...")
		fyne.NewAnimation(4*time.Second, doAnimation).Start()
	})

	return container.NewVBox(
		container.NewHBox(
			layout.NewSpacer(),
			img[left],
			img[right],
			layout.NewSpacer(),
		),
		rollButton,
		result,
	)
}
