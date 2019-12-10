package main

import (
	"fmt"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"path/filepath"
)

func main() {

	rect := sciter.NewRect(600, 600, 600, 600)
	window, windowCreationErr := window.New(
		sciter.SW_MAIN|
		sciter.SW_CONTROLS|
		sciter.SW_RESIZEABLE|
		// sciter.SW_ALPHA|  // removes all borders
		sciter.SW_ENABLE_DEBUG, rect)

	if windowCreationErr != nil {
		fmt.Errorf("Could not create sciter window : %s",
			windowCreationErr.Error())
		return
	}

	var filePath, _ = filepath.Abs("./index.html")
	uiFileLoadErr := window.LoadFile(filePath)
	if uiFileLoadErr != nil {
		fmt.Errorf("Could not load ui file : %s",
			uiFileLoadErr.Error())
	}

	// Setting up stage for Harmony
	window.SetTitle("Simple Input")
	window.Show()
	window.Run()

}