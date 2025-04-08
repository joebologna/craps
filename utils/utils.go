package utils

import (
	"fyne.io/fyne/v2"
)

func GetDeviceType() (is_mobile, is_browser, is_desktop bool) {
	is_mobile, is_browser = fyne.CurrentDevice().IsMobile(), fyne.CurrentDevice().IsBrowser()
	is_desktop = !(is_mobile || is_browser)
	return
}

func GetScreenSize() fyne.Size {
	// is_mobile, is_browser, _ := GetDeviceType()
	// o := fyne.CurrentDevice().Orientation()
	// if is_mobile || is_browser {
	// 	if o == fyne.OrientationVertical || o == fyne.OrientationVerticalUpsideDown {
	// 		return fyne.NewSize(768, 1024)
	// 	}
	// 	return fyne.NewSize(1024, 768)
	// }
	// if o == fyne.OrientationVertical || o == fyne.OrientationVerticalUpsideDown {
	// 	return fyne.NewSize(768, 1024)
	// }
	return fyne.NewSquareSize(256)
}
