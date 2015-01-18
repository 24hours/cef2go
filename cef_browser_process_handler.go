// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cef2go

package chrome

/*
#include <stdlib.h>
#include "cef_base.h"
#include "include/capi/cef_browser_process_handler_capi.h"
extern void intialize_cef_browser_process_handler(struct _cef_browser_process_handler_t* handler);
*/
import "C"

import "unsafe"

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

var (
	browserProcessHandlerMap = make(map[unsafe.Pointer]BrowserProcessHandler)
)

type BrowserProcessHandler interface {
	///
	// Called on the browser process UI thread immediately after the CEF context
	// has been initialized.
	///
	OnContextInitialized()

	///
	// Called before a child process is launched. Will be called on the browser
	// process UI thread when launching a render process and on the browser
	// process IO thread when launching a GPU or plugin process. Provides an
	// opportunity to modify the child process command line. Do not keep a
	// reference to |command_line| outside of this function.
	///
	OnBeforeChildProcessLaunch(commandLine CommandLineT)

	///
	// Called on the browser process IO thread after the main thread has been
	// created for a new render process. Provides an opportunity to specify extra
	// information that will be passed to
	// cef_render_process_handler_t::on_render_thread_created() in the render
	// process. Do not keep a reference to |extra_info| outside of this function.
	///
	OnRenderProcessThreadCreated(extraInfo CefListValueT)

	// returns the browser process handler T struct
	GetBrowserProcessHandlerT() BrowserProcessHandlerT
}

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
