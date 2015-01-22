// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cef2go

package chrome

/*
#cgo CFLAGS: -I./
#include "cef_base.h"
#include "include/capi/cef_browser_process_handler_capi.h"
*/
import "C"
import "unsafe"

//BaseBrowserProcessHandler
type BaseBrowserProcessHandler struct {
	handler BrowserProcessHandlerT
}

func (bph *BaseBrowserProcessHandler) OnContextInitialized() {}
func (bph *BaseBrowserProcessHandler) OnBeforeChildProcessLaunch(commandLine CommandLineT) {
	defer commandLine.Release()
}

func (bph *BaseBrowserProcessHandler) OnRenderProcessThreadCreated(extraInfo CefListValueT) {
	defer extraInfo.Release()
}

func (bph *BaseBrowserProcessHandler) GetBrowserProcessHandlerT() BrowserProcessHandlerT {
	return bph.handler
}

type CommandLineT struct {
	CStruct *C.struct__cef_command_line_t
}

func (c CommandLineT) Release() {
	Release(unsafe.Pointer(c.CStruct))
}

type CefListValueT struct {
	CStruct *C.struct__cef_list_value_t
}

func (c CefListValueT) AddRef() {
	AddRef(unsafe.Pointer(c.CStruct))
}
func (c CefListValueT) Release() {
	Release(unsafe.Pointer(c.CStruct))
}

type BrowserProcessHandlerT struct {
	CStruct *C.struct__cef_browser_process_handler_t
}

func (c BrowserProcessHandlerT) AddRef() {
	AddRef(unsafe.Pointer(c.CStruct))
}
func (c BrowserProcessHandlerT) Release() {
	Release(unsafe.Pointer(c.CStruct))
}
