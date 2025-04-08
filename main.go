package main

import (
	"bytes"
	"craps/utils"
	"embed"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//go:embed media/Animation/*
var animationFiles embed.FS

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&CustomTheme{theme.DefaultTheme()})
	myWindow := myApp.NewWindow("Craps")

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
	rollButton = widget.NewButton("Roll", func() {
		if !inProgress {
			showImages(images, &inProgress, rollButton, &resultString)
		}
	})
	images[len(images)-1].Show()
	myWindow.SetContent(
		container.NewVBox(
			container.NewCenter(container.NewStack(images...)),
			rollButton,
			result,
		),
	)
	screenSize := utils.GetScreenSize()
	myWindow.Resize(screenSize)

	myWindow.ShowAndRun()
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
