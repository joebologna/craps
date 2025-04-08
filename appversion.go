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
	AppVersion2a
	AppVersion2b
	AppVersion2c
	AppVersion3
)

func (app AppVersion) String() string {
	switch app {
	case AppVersion1:
		return "V1"
	case AppVersion2a:
		return "V2a"
	case AppVersion2b:
		return "V2b"
	case AppVersion2c:
		return "V2c"
	default:
		return "Unknown"
	}
}

func (v AppVersion) App(animationFiles embed.FS) (stuff *fyne.Container) {
	switch v {
	case AppVersion1:
		return exp.App1(animationFiles)
	case AppVersion2a:
		return exp.App2(animationFiles, opts.GoFuncOpt)
	case AppVersion2b:
		return exp.App2(animationFiles, opts.AnimationWithShowHide)
	case AppVersion2c:
		return exp.App2(animationFiles, opts.AnimateImageObject)
	default:
		panic("unsupported version")
	}
}
