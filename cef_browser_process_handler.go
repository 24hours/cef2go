// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cef2go

package chrome

/*
#cgo CFLAGS: -I./
#include <stdlib.h>
#include "cef_base.h"
#include "include/capi/cef_browser_process_handler_capi.h"
extern void intialize_cef_browser_process_handler(struct _cef_browser_process_handler_t* handler);
*/
import "C"
import "unsafe"

var browserProcessHandlerMap = make(map[unsafe.Pointer]BrowserProcessHandler)

//export OnContextInitialized
func OnContextInitialized(self *C.struct__cef_browser_process_handler_t) {
	if handler, ok := browserProcessHandlerMap[unsafe.Pointer(self)]; ok {
		handler.OnContextInitialized()
	}
}

//export OnBeforeChildProcessLaunch
func OnBeforeChildProcessLaunch(self *C.struct__cef_browser_process_handler_t, commandLine *C.struct__cef_command_line_t) {
	if handler, ok := browserProcessHandlerMap[unsafe.Pointer(self)]; ok {
		handler.OnBeforeChildProcessLaunch(CommandLineT{commandLine})
		return
	}
	CommandLineT{commandLine}.Release()
}

//export OnRenderProcessThreadCreated
func OnRenderProcessThreadCreated(self *C.struct__cef_browser_process_handler_t, extraInfo *C.struct__cef_list_value_t) {
	if handler, ok := browserProcessHandlerMap[unsafe.Pointer(self)]; ok {
		handler.OnRenderProcessThreadCreated(CefListValueT{extraInfo})
		return
	}
	CefListValueT{extraInfo}.Release()
}

func NewBrowserProcessHandlerT(handler BrowserProcessHandler) BrowserProcessHandlerT {
	var b BrowserProcessHandlerT
	b.CStruct = (*C.struct__cef_browser_process_handler_t)(
		C.calloc(1, C.sizeof_struct__cef_browser_process_handler_t))
	C.intialize_cef_browser_process_handler(b.CStruct)
	unsafeIt := unsafe.Pointer(b.CStruct)
	go_AddRef(unsafeIt)
	browserProcessHandlerMap[unsafeIt] = handler
	return b
}

type BrowserProcessHandler interface {
	OnContextInitialized()
	OnBeforeChildProcessLaunch(commandLine CommandLineT)
	OnRenderProcessThreadCreated(extraInfo CefListValueT)
	GetBrowserProcessHandlerT() BrowserProcessHandlerT
}
