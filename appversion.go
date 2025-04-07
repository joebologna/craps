package main

import (
	"craps/apps/exp"

	"fyne.io/fyne/v2"
)

type AppVersion int

const (
	AppVersion1 AppVersion = iota
)

func (app AppVersion) String() string {
	switch app {
	case AppVersion1:
		return "V1"
	default:
		return "Unknown"
	}
}

func (v AppVersion) App() (stuff *fyne.Container, button fyne.CanvasObject) {
	switch v {
	case AppVersion1:
		return exp.App1()
	default:
		panic("unsupported version")
	}
}
