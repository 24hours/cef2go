// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cefcapi

package chrome

/*
#include <stdlib.h>
#include "cef_base.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_life_span_handler_capi.h"
extern void initialize_life_span_handler(struct _cef_life_span_handler_t* lifeHandler);
*/
import "C"
import "unsafe"

var lifeSpanHandlerMap = make(map[unsafe.Pointer]LifeSpanHandler)

type LifeSpanHandler interface {
	OnAfterCreated(browser CefBrowserT)
	RunModal(browser CefBrowserT) int
	DoClose(browser CefBrowserT) int
	BeforeClose(browser CefBrowserT)

	GetLifeSpanHandlerT() LifeSpanHandlerT
}

type LifeSpanHandlerT struct {
	CStruct *C.struct__cef_life_span_handler_t
}

func (r LifeSpanHandlerT) AddRef() {
	AddRef(unsafe.Pointer(r.CStruct))
}
func (r LifeSpanHandlerT) Release() {
	Release(unsafe.Pointer(r.CStruct))
}

//export go_OnBeforePopup
func go_OnBeforePopup(
	self *C.struct__cef_life_span_handler_t,
	browser *C.struct__cef_browser_t,
	frame *C.struct__cef_frame_t,
	target_url *C.cef_string_t,
	target_frame_name *C.cef_string_t,
	popupFeatures *C.struct__cef_popup_features_t,
	windowInfo *C.struct__cef_window_info_t,
	client **C.struct__cef_client_t,
	settings *C.struct__cef_browser_settings_t,
	no_javascript_access *C.int) int {

	C.releaseVoid(unsafe.Pointer(browser))
	C.releaseVoid(unsafe.Pointer(frame))
	C.releaseVoid(unsafe.Pointer(popupFeatures))
	C.releaseVoid(unsafe.Pointer(windowInfo))
	//C.releaseVoid(unsafe.Pointer(client))
	C.releaseVoid(unsafe.Pointer(settings))
	return 1
}

//export go_OnAfterCreated
func go_OnAfterCreated(
	self *C.struct__cef_life_span_handler_t,
	browser *C.struct__cef_browser_t) {

	b := CefBrowserT{browser}
	if handler, ok := lifeSpanHandlerMap[unsafe.Pointer(self)]; ok {
		handler.OnAfterCreated(b)
		return
	}
	b.Release()
}

//export go_RunModal
func go_RunModal(
	self *C.struct__cef_life_span_handler_t,
	browser *C.struct__cef_browser_t) int {
	b := CefBrowserT{browser}
	if handler, ok := lifeSpanHandlerMap[unsafe.Pointer(self)]; ok {
		return handler.RunModal(b)
	}
	b.Release()
	return 0
}

//export go_DoClose
func go_DoClose(
	self *C.struct__cef_life_span_handler_t,
	browser *C.struct__cef_browser_t) int {

	b := CefBrowserT{browser}
	if handler, ok := lifeSpanHandlerMap[unsafe.Pointer(self)]; ok {
		return handler.DoClose(b)
	}
	b.Release()
	return 0
}

//export go_BeforeClose
func go_BeforeClose(
	self *C.struct__cef_life_span_handler_t,
	browser *C.struct__cef_browser_t) {
	b := CefBrowserT{browser}
	if handler, ok := lifeSpanHandlerMap[unsafe.Pointer(self)]; ok {
		handler.BeforeClose(b)
		return
	}
	b.Release()
}

func NewLifeSpanHandlerT(life LifeSpanHandler) LifeSpanHandlerT {
	var handler LifeSpanHandlerT
	handler.CStruct = (*C.struct__cef_life_span_handler_t)(
		C.calloc(1, C.sizeof_struct__cef_life_span_handler_t))
	C.initialize_life_span_handler(handler.CStruct)
	go_AddRef(unsafe.Pointer(handler.CStruct))
	lifeSpanHandlerMap[unsafe.Pointer(handler.CStruct)] = life
	return handler
}

//base LifeSpanHandler

type BaseLifeSpanHandler struct {
	lifeSpan LifeSpanHandlerT
}

func (l *BaseLifeSpanHandler) OnAfterCreated(browser CefBrowserT) {
	defer browser.Release()

}
func (l *BaseLifeSpanHandler) RunModal(browser CefBrowserT) int {
	defer browser.Release()

	return 0
}
func (l *BaseLifeSpanHandler) DoClose(browser CefBrowserT) int {
	defer browser.Release()

	return 0
}
func (l *BaseLifeSpanHandler) BeforeClose(browser CefBrowserT) {
	defer browser.Release()

	//chrome.QuitMessageLoop()
}
func (l *BaseLifeSpanHandler) GetLifeSpanHandlerT() LifeSpanHandlerT {
	return l.lifeSpan
}
