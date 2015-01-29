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
	chrome.CreateBrowser(window, ch, chrome.BrowserSettings{}, "http://www.imgur.com")
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

func (l *myClientHandler) OnAfterCreated(browser chrome.Browser) {
	defer browser.Release()
	fmt.Println("lifespan::OnAfterCreated")
}
func (l *myClientHandler) RunModal(browser chrome.Browser) int {
	fmt.Println("lifespan::RunModal")
	return 0
}
func (l *myClientHandler) DoClose(browser chrome.Browser) int {
	fmt.Println("lifespan::DoClose")
	return 0
}
func (l *myClientHandler) BeforeClose(browser chrome.Browser) {
	fmt.Println("lifespan::BeforeClose")
}

func (l *myClientHandler) OnBeforeBrowse(browser chrome.Browser, frame chrome.CefFrameT, request chrome.CefRequestT, isRedirect int) int {
	fmt.Println("Before browse: ", request.GetUrl())
	return 0
}
func (l *myClientHandler) OnBeforeResourceLoad(browser chrome.Browser, frame chrome.CefFrameT, request chrome.CefRequestT) int {
	fmt.Println("Resource Browse", request.GetUrl())
	return 0
}
func (l *myClientHandler) OnCertificateError(errorCode chrome.CefErrorCode, requestUrl string, errorCallback chrome.CefCertErrorCallbackT) int {
	return 0
}

func (d *myClientHandler) OnAddressChange(browser chrome.Browser, frame chrome.CefFrameT, url string) {

}

func (d *myClientHandler) OnTitleChange(browser chrome.Browser, title string) {

}

func (d *myClientHandler) OnToolTip(browser chrome.Browser, text string) bool {
	fmt.Println("Tooltip: ", text)
	return true
}

func (d *myClientHandler) OnStatusMessage(browser chrome.Browser, value string) {
	fmt.Println("Status: ", value)
}

func (d *myClientHandler) OnConsoleMessage(browser chrome.Browser, message, source string, line int) bool {
	fmt.Println("Console:[", source, ":", line, "] ", message)
	return true
}
