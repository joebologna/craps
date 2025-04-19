Run:
    make app_store

The output will show where the working directory is. Set WORK environment variable from the output, for example:

    WORK=/var/folders/3f/5d8mfhnj5sn2r7kr8711h_jc0000gn/T/fyne-work-4285550681

Copy the project that includes the new binary:

    rsync -av $WORK/main $WORK/arm64 fyne-work/

Verify the version string is correct:

    plutil -p fyne-work/main/Info.plist|grep ShortVersionString

Open the project and verify the images are set properly.

    open fyne-work/main.xcodeproj

Run Archive, Validate, Distribute.

Wait for the build to complete, select in in App Store Connect.
