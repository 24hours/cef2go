// Copyright (c) 2014 The chrome authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/chrome

package main

import (
	"github.com/24hours/chrome"
	"github.com/24hours/chrome/ui"
	"os"
)

func main() {
	chrome.ExecuteProcess(nil, nil)

	// Initialize CEF.
	settings := chrome.NewSettings()
	settings.LocalesDirPath = ""
	settings.ResourcesDirPath = ""
	chrome.Initialize(settings, nil)

	// Create Window using Cocoa API.
	window := ui.CreateWindow("chrome example", 1024, 768)
	ui.ConnectDestroySignal(window, OnDestroyWindow)
	ui.ActivateApp()

	chrome.CreateBrowser(window, nil, chrome.BrowserSettings{}, "http://www.google.com")
	// CEF loop and shutdown.
	chrome.RunMessageLoop()
	//chrome.Shutdown()
	os.Exit(0)
}

func OnDestroyWindow() {
	chrome.QuitMessageLoop()
}
