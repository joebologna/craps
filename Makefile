sim:
	fyne package -os iossimulator -certificate "Apple Development: Joe Bologna (Q5BFX4NKWR)" -profile "Craps"
	-xcrun simctl boot "iPhone 16 Pro Max"
	xcrun simctl install "iPhone 16 Pro Max" craps.app

phone:
	fyne package -os ios -certificate "Apple Development: Joe Bologna (Q5BFX4NKWR)" -profile "Craps"

# this is unused
dist:
	fyne package -os ios -certificate "Apple Distribution: Focused for Success, Inc. (2GC862GT48)" -profile "Craps App Store"

app_store:
	../fyne/cmd/fyne/fyne package -work -os ios -certificate "Apple Distribution: Focused for Success, Inc. (2GC862GT48)" -profile "Craps App Store" -release
	@echo open Xcode project, run gen_icons.sh, install the icons, update the project to use the Craps App Store provisioning profile, run Archive
