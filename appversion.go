package main

import (
	"craps/apps/exp"
	"craps/opts"
	"embed"

	"fyne.io/fyne/v2"
)

type AppVersion int

const (
	AppVersion1 AppVersion = iota
	AppVersion2
)

func (app AppVersion) String() string {
	switch app {
	case AppVersion1:
		return "V1"
	case AppVersion2:
		return "V2"
	default:
		return "Unknown"
	}
}

func (v AppVersion) App(animationFiles embed.FS) (stuff *fyne.Container) {
	switch v {
	case AppVersion1:
		return exp.App1(animationFiles)
	case AppVersion2:
		return exp.App2(animationFiles, opts.AnimateImageObject)
	default:
		panic("unsupported version")
	}
}
