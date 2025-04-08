package main

import (
	"craps/utils"
	"embed"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

//go:embed media/Animation/*
var animationFiles embed.FS

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&CustomTheme{theme.DefaultTheme()})
	myWindow := myApp.NewWindow("Craps")

	screenSize := utils.GetScreenSize()
	myWindow.Resize(screenSize)

	stuff := AppVersion2c.App(animationFiles)

	myWindow.SetContent(stuff)

	myWindow.ShowAndRun()
}
