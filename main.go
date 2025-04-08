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
	"fyne.io/fyne/v2/theme"
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

	myWindow.SetContent(container.NewCenter(images...))
	screenSize := utils.GetScreenSize()
	myWindow.Resize(screenSize)

	go func() {
		images[0].Show()
		time.Sleep(5 * time.Second)
		var i = 0
		for i = 0; i < len(images); i++ {
			images[i].Show()
			time.Sleep(time.Millisecond * 150)
			images[i].Hide()
		}
		images[len(images)-1].Show()
	}()

	// Example usage of animationFiles
	// You can load files from the embedded filesystem as needed
	// Example: data, _ := animationFiles.ReadFile("media/Animation/example.gif")

	myWindow.ShowAndRun()
}
