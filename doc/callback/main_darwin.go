// Copyright (c) 2014 The chrome authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/chrome

package main

import (
	"fmt"
	"github.com/24hours/chrome"
	"github.com/24hours/chrome/cocoa"
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
	window := cocoa.CreateWindow("chrome example", 1024, 768)
	cocoa.ConnectDestroySignal(window, OnDestroyWindow)
	cocoa.ActivateApp()

	ch := NewMyClientHandler()
	chrome.CreateBrowser(window, ch, chrome.BrowserSettings{}, "http://www.google.com")
	// CEF loop and shutdown.
	chrome.RunMessageLoop()
	//chrome.Shutdown()
	os.Exit(0)
}

func OnDestroyWindow() {
	chrome.QuitMessageLoop()
}

func NewMyClientHandler() *myClientHandler {
	ch := &myClientHandler{}

	return ch
}

type myClientHandler struct {
	chrome.BaseClientHandler
}

func (l *myClientHandler) OnAfterCreated(browser chrome.CefBrowserT) {
	defer browser.Release()
	fmt.Println("lifespan::OnAfterCreated")
}
func (l *myClientHandler) RunModal(browser chrome.CefBrowserT) int {
	defer browser.Release()
	fmt.Println("lifespan::RunModal")
	return 0
}
func (l *myClientHandler) DoClose(browser chrome.CefBrowserT) int {
	defer browser.Release()
	fmt.Println("lifespan::DoClose")
	return 0
}
func (l *myClientHandler) BeforeClose(browser chrome.CefBrowserT) {
	defer browser.Release()
	fmt.Println("lifespan::BeforeClose")
	chrome.QuitMessageLoop()
}
