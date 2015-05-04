// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package main

import (
	"fmt"
	"github.com/24hours/chrome"
	"github.com/24hours/chrome/ui"
	"os"
)

func main() {
	chrome.ExecuteProcess(nil, nil)
	setting := chrome.NewSettings()
	setting.NoSandbox = 0
	setting.SingleProcess = 0
	chrome.Initialize(setting, nil)
	fmt.Println("create")
	window := ui.CreateWindow("chrome example", 1024, 768)
	ui.ConnectDestroySignal(window, OnDestroyWindow)

	fmt.Println("create browser")
	// Create browser.
	chrome.CreateBrowser(window, nil, chrome.BrowserSettings{}, "http://www.google.com")

	// CEF loop and shutdown.
	chrome.RunMessageLoop()
	//chrome.Shutdown()
	os.Exit(0)
}

func OnDestroyWindow() {
	chrome.QuitMessageLoop()
}
