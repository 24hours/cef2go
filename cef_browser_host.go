// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cefcapi

package chrome

/*
#include <stdlib.h>
#include "cef_base.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_browser_capi.h"
extern struct _cef_browser_t* cef_browser_host_t_get_browser(struct _cef_browser_host_t* self);
extern void cef_browser_host_t_close_browser(struct _cef_browser_host_t* self,int force_close);
extern void cef_browser_host_t_set_focus(struct _cef_browser_host_t* self, int enable);
extern cef_window_handle_t cef_browser_host_t_get_window_handle(struct _cef_browser_host_t* self);
extern cef_window_handle_t cef_browser_host_t_get_opener_window_handle(struct _cef_browser_host_t* self);
extern struct _cef_client_t* cef_browser_host_t_get_client(struct _cef_browser_host_t* self);
extern struct _cef_request_context_t* cef_browser_host_t_get_request_context(struct _cef_browser_host_t* self);
//extern cef_string_utf8_t* cef_browser_host_t_get_dev_tools_url(struct _cef_browser_host_t* self, int http_scheme);
extern double cef_browser_host_t_get_zoom_level(struct _cef_browser_host_t* self);
extern void cef_browser_host_t_set_zoom_level(struct _cef_browser_host_t* self, double zoomLevel);
extern void cef_browser_host_t_run_file_dialog(struct _cef_browser_host_t* self, cef_file_dialog_mode_t mode, const cef_string_t* title,const cef_string_t* default_file_name, cef_string_list_t accept_types,struct _cef_run_file_dialog_callback_t* callback);
extern void cef_browser_host_t_start_download(struct _cef_browser_host_t* self,const cef_string_t* url);
extern void cef_browser_host_t_print(struct _cef_browser_host_t* self);
extern void cef_browser_host_t_find(struct _cef_browser_host_t* self, int identifier,const cef_string_t* searchText, int forward, int matchCase,int findNext);
extern void cef_browser_host_t_stop_finding(struct _cef_browser_host_t* self,int clearSelection);
extern void cef_browser_host_t_set_mouse_cursor_change_disabled(struct _cef_browser_host_t* self, int disabled);
extern int cef_browser_host_t_is_mouse_cursor_change_disabled(struct _cef_browser_host_t* self);
extern int cef_browser_host_t_is_window_rendering_disabled(struct _cef_browser_host_t* self);
extern void cef_browser_host_t_was_resized(struct _cef_browser_host_t* self);
extern void cef_browser_host_t_was_hidden(struct _cef_browser_host_t* self, int hidden);
extern void cef_browser_host_t_notify_screen_info_changed(struct _cef_browser_host_t* self);
extern void cef_browser_host_t_invalidate(struct _cef_browser_host_t* self, cef_paint_element_type_t type);
extern void cef_browser_host_t_send_key_event(struct _cef_browser_host_t* self, const struct _cef_key_event_t* event);
extern void cef_browser_host_t_send_mouse_click_event(struct _cef_browser_host_t* self, const struct _cef_mouse_event_t* event, cef_mouse_button_type_t type, int mouseUp, int clickCount);
extern void cef_browser_host_t_send_mouse_move_event(struct _cef_browser_host_t* self,const struct _cef_mouse_event_t* event, int mouseLeave);
extern void cef_browser_host_t_send_mouse_wheel_event(struct _cef_browser_host_t* self,const struct _cef_mouse_event_t* event, int deltaX, int deltaY);
extern void cef_browser_host_t_send_focus_event(struct _cef_browser_host_t* self,int setFocus);
extern void cef_browser_host_t_send_capture_lost_event(struct _cef_browser_host_t* self);
extern cef_text_input_context_t cef_browser_host_t_get_nstext_input_context(struct _cef_browser_host_t* self);
extern void cef_browser_host_t_handle_key_event_before_text_input_client(struct _cef_browser_host_t* self, cef_event_handle_t keyEvent);
extern void cef_browser_host_t_handle_key_event_after_text_input_client(struct _cef_browser_host_t* self, cef_event_handle_t keyEvent);
*/
import "C"

import (
	"unsafe"
)

type CefBrowserHostT struct {
	CStruct *C.struct__cef_browser_host_t
}

func (c CefBrowserHostT) Release() {
	Release(unsafe.Pointer(c.CStruct))
}
func (c CefBrowserHostT) AddRef() {
	AddRef(unsafe.Pointer(c.CStruct))
}

type CefRequestContextT struct {
	CStruct *C.struct__cef_request_context_t
}
type CefRunFileDialogCallbackT struct {
	CStruct *C.struct__cef_run_file_dialog_callback_t
}
type CefMouseEventT struct {
	CStruct *C.struct__cef_mouse_event_t
}
type CefKeyEventT struct {
	CStruct *C.struct__cef_key_event_t
}
type CefEventHandleT struct {
	// CStruct *C.MSG
}
type CefClientT struct {
	CStruct *C.struct__cef_client_t
}
type CefRectT struct {
	CStruct *C.cef_rect_t
}

func (b CefBrowserHostT) GetBrowser() CefBrowserT {
	return CefBrowserT{C.cef_browser_host_t_get_browser(b.CStruct)}
}

func (b CefBrowserHostT) CloseBrowser(forceClose int) {
	C.cef_browser_host_t_close_browser(b.CStruct, C.int(forceClose))
}

func (b CefBrowserHostT) SetFocus(enable int) {
	C.cef_browser_host_t_set_focus(b.CStruct, C.int(enable))
}

func (b CefBrowserHostT) GetWindowHandle() C.cef_window_handle_t {
	return C.cef_browser_host_t_get_window_handle(b.CStruct)
}

func (b CefBrowserHostT) GetOpenerWindowHandle() C.cef_window_handle_t {
	return C.cef_browser_host_t_get_opener_window_handle(b.CStruct)
}

func (b CefBrowserHostT) GetClient() CefClientT {
	return CefClientT{C.cef_browser_host_t_get_client(b.CStruct)}
}

func (b CefBrowserHostT) GetRequestContext() CefRequestContextT {
	return CefRequestContextT{C.cef_browser_host_t_get_request_context(b.CStruct)}
}

func (b CefBrowserHostT) GetZoomLevel() float64 {
	return float64(C.cef_browser_host_t_get_zoom_level(b.CStruct))
}

func (b CefBrowserHostT) SetZoomLevel(zoomLevel float64) {
	C.cef_browser_host_t_set_zoom_level(b.CStruct, C.double(zoomLevel))
}

func (b CefBrowserHostT) RunFileDialog(mode int, title string, defaultFileName string, acceptTypes []string, callback CefRunFileDialogCallbackT) {
	//C.cef_browser_host_t_run_file_dialog(b.CStruct)
}

func (b CefBrowserHostT) StartDownload(url string) {
	//C.cef_browser_host_t_start_download(b.CStruct)
}

func (b CefBrowserHostT) Print() {
	C.cef_browser_host_t_print(b.CStruct)
}

func (b CefBrowserHostT) Find(identifier int, search string, forward int, matchCase int, findNext int) {
	//C.cef_browser_host_t_find(b.CStruct)
}

func (b CefBrowserHostT) StopFinding(clearSelection int) {
	C.cef_browser_host_t_stop_finding(b.CStruct, C.int(clearSelection))
}

func (b CefBrowserHostT) SetMouseCursorChangeDisabled(disabled int) {
	C.cef_browser_host_t_set_mouse_cursor_change_disabled(b.CStruct, C.int(disabled))
}

func (b CefBrowserHostT) IsMouseCursorChangeDisabled() int {
	return int(C.cef_browser_host_t_is_mouse_cursor_change_disabled(b.CStruct))
}

func (b CefBrowserHostT) IsWindowRenderingDisabled() int {
	return int(C.cef_browser_host_t_is_window_rendering_disabled(b.CStruct))
}

func (b CefBrowserHostT) WasResized() {
	C.cef_browser_host_t_was_resized(b.CStruct)
}

func (b CefBrowserHostT) WasHidden(hidden int) {
	C.cef_browser_host_t_was_hidden(b.CStruct, C.int(hidden))
}

func (b CefBrowserHostT) NotifyScreenInfoChange() {
	C.cef_browser_host_t_notify_screen_info_changed(b.CStruct)
}

func (b CefBrowserHostT) Invalidate(paintType C.cef_paint_element_type_t) {
	C.cef_browser_host_t_invalidate(b.CStruct, paintType)
}

func (b CefBrowserHostT) SendKeyEvent(event CefKeyEventT) {
	C.cef_browser_host_t_send_key_event(b.CStruct, event.CStruct)
}

func (b CefBrowserHostT) SendMouseClickEvent(event CefMouseEventT, buttonType C.cef_mouse_button_type_t, mouseUp, clickCount int) {
	C.cef_browser_host_t_send_mouse_click_event(b.CStruct, event.CStruct, buttonType, C.int(mouseUp), C.int(clickCount))
}

func (b CefBrowserHostT) SendMouseMoveEvent(event CefMouseEventT, mouseLeave int) {
	C.cef_browser_host_t_send_mouse_move_event(b.CStruct, event.CStruct, C.int(mouseLeave))
}

func (b CefBrowserHostT) SendMouseWheelEVent(event CefMouseEventT, deltaX, deltaY int) {
	C.cef_browser_host_t_send_mouse_wheel_event(b.CStruct, event.CStruct, C.int(deltaX), C.int(deltaY))
}

func (b CefBrowserHostT) SendFocusEvent(setFocus int) {
	C.cef_browser_host_t_send_focus_event(b.CStruct, C.int(setFocus))
}

func (b CefBrowserHostT) SendCaptureLostEvent() {
	C.cef_browser_host_t_send_capture_lost_event(b.CStruct)
}

func (b CefBrowserHostT) GetNstextInputContext() { //return type?
	C.cef_browser_host_t_get_nstext_input_context(b.CStruct)
}

// func (b CefBrowserHostT) HandleKeyEventBeforeTextInputCLient(keyEvent CefEventHandleT) {
// 	C.cef_browser_host_t_handle_key_event_before_text_input_client(b.CStruct, keyEvent.CStruct)
// }

// func (b CefBrowserHostT) HandleKeyEventAfterTextInputCLient(keyEvent CefEventHandleT) {
// 	C.cef_browser_host_t_handle_key_event_after_text_input_client(b.CStruct, keyEvent.CStruct)
// }
