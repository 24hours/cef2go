// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cef2go

package chrome

/*
#include <stdlib.h>
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_display_handler_capi.h"
extern void initialize_display_handler(struct _cef_display_handler_t * displayHandler);
*/
import "C"
import (
	"unsafe"
)

type DisplayHandler interface {
	OnAddressChange(browser Browser, frame CefFrameT, url string)
	OnTitleChange(browser Browser, title string)
	OnToolTip(browser Browser, text string) bool
	OnStatusMessage(browser Browser, value string)
	OnConsoleMessage(browser Browser, message, source string, line int) bool

	GetDisplayHandler() DisplayHandlerT
}

var displayHandlerMap = make(map[unsafe.Pointer]DisplayHandler)

type DisplayHandlerT struct {
	CStruct *C.struct__cef_display_handler_t
}

func (r DisplayHandlerT) AddRef() {
	AddRef(unsafe.Pointer(r.CStruct))
}
func (r DisplayHandlerT) Release() {
	Release(unsafe.Pointer(r.CStruct))
}

//export go_OnAddressChange
func go_OnAddressChange(self *C.struct__cef_display_handler_t,
	browser *C.struct__cef_browser_t,
	frame *C.struct__cef_frame_t,
	url *C.char) {
	defer Browser{browser}.Release()
	defer CefFrameT{frame}.Release()

	if handler, ok := displayHandlerMap[unsafe.Pointer(self)]; ok {
		handler.OnAddressChange(
			Browser{browser},
			CefFrameT{frame},
			C.GoString(url),
		)
		return
	}
}

//export go_OnTitleChange
func go_OnTitleChange(self *C.struct__cef_display_handler_t,
	browser *C.struct__cef_browser_t,
	title *C.char) {
	defer Browser{browser}.Release()
	if handler, ok := displayHandlerMap[unsafe.Pointer(self)]; ok {
		handler.OnTitleChange(
			Browser{browser},
			C.GoString(title),
		)
		return
	}

}

//export go_OnTooltip
func go_OnTooltip(self *C.struct__cef_display_handler_t,
	browser *C.struct__cef_browser_t,
	text *C.char) int {
	defer Browser{browser}.Release()
	if handler, ok := displayHandlerMap[unsafe.Pointer(self)]; ok {
		if bVal := handler.OnToolTip(
			Browser{browser},
			C.GoString(text),
		); bVal {
			return 1
		}
	}
	return 0

}

//export go_OnStatusMessage
func go_OnStatusMessage(self *C.struct__cef_display_handler_t,
	browser *C.struct__cef_browser_t,
	value *C.char) {
	defer Browser{browser}.Release()
	if handler, ok := displayHandlerMap[unsafe.Pointer(self)]; ok {
		handler.OnStatusMessage(
			Browser{browser},
			C.GoString(value),
		)
		return
	}
}

//export go_OnConsoleMessage
func go_OnConsoleMessage(self *C.struct__cef_display_handler_t,
	browser *C.struct__cef_browser_t,
	message *C.char,
	source *C.char,
	line C.int) int {
	defer Browser{browser}.Release()
	if handler, ok := displayHandlerMap[unsafe.Pointer(self)]; ok {
		if bVal := handler.OnConsoleMessage(
			Browser{browser},
			C.GoString(message),
			C.GoString(source),
			int(line),
		); bVal {
			return 1
		}
	}
	return 0
}

func NewDisplayHandlerT(display DisplayHandler) DisplayHandlerT {
	var handler DisplayHandlerT
	handler.CStruct = (*C.struct__cef_display_handler_t)(
		C.calloc(1, C.sizeof_struct__cef_display_handler_t))
	C.initialize_display_handler(handler.CStruct)
	go_AddRef(unsafe.Pointer(handler.CStruct))
	displayHandlerMap[unsafe.Pointer(handler.CStruct)] = display
	return handler
}
