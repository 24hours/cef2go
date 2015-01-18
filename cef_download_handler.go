// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cef2go

package chrome

/*
#include <stdlib.h>
#include "cef_base.h"
#include "include/capi/cef_client_capi.h"
extern void initialize_download_handler(struct _cef_download_handler_t* self);
extern void callDownloadItemCallback_cancel(struct _cef_download_item_callback_t* self);
extern void callBeforeDownloadCallback_cont(struct _cef_before_download_callback_t* self, const char * download_path, int show_dialog);
extern int callCefDownloadItem_is_valid(struct _cef_download_item_t* self);
extern int callCefDownloadItem_is_in_progress(struct _cef_download_item_t* self);
extern int callCefDownloadItem_is_complete(struct _cef_download_item_t* self);
extern int callCefDownloadItem_is_canceled(struct _cef_download_item_t* self);
extern int64 callCefDownloadItem_get_current_speed(struct _cef_download_item_t* self);
extern int callCefDownloadItem_get_percent_complete(struct _cef_download_item_t* self);
extern int64 callCefDownloadItem_get_total_bytes(struct _cef_download_item_t* self);
extern int64 callCefDownloadItem_get_received_bytes(struct _cef_download_item_t* self);
extern cef_time_t callCefDownloadItem_get_start_time(struct _cef_download_item_t* self);
extern cef_time_t callCefDownloadItem_get_end_time(struct _cef_download_item_t* self);
extern cef_string_utf8_t* callCefDownloadItem_get_full_path(struct _cef_download_item_t* self);
extern uint32 callCefDownloadItem_get_id(struct _cef_download_item_t* self);
extern cef_string_utf8_t* callCefDownloadItem_get_url(struct _cef_download_item_t* self);
extern cef_string_utf8_t* callCefDownloadItem_get_suggested_file_name(struct _cef_download_item_t* self);
extern cef_string_utf8_t* callCefDownloadItem_get_content_disposition(struct _cef_download_item_t* self);
extern cef_string_utf8_t* callCefDownloadItem_get_mime_type(struct _cef_download_item_t* self);
*/
import "C"
import "unsafe"

var (
	downloadHandlerMap = make(map[unsafe.Pointer]DownloadHandler)
)

type DownloadHandler interface {
	OnBeforeDownload(browser CefBrowserT, downloadItem CefDownloadItemT, suggestedName string, callback CefBeforeDownloadCallbackT) string
	OnDownloadUpdated(browser CefBrowserT, downloadItem CefDownloadItemT, callback CefDownloadItemCallbackT)

	GetDownloadHandlerT() DownloadHandlerT
}

type DownloadHandlerT struct {
	CStruct *C.struct__cef_download_handler_t
}

func (r DownloadHandlerT) AddRef() {
	AddRef(unsafe.Pointer(r.CStruct))
}
func (r DownloadHandlerT) Release() {
	Release(unsafe.Pointer(r.CStruct))
}

type CefTimeT struct {
	Self C.cef_time_t
}

type CefBeforeDownloadCallbackT struct {
	Self *C.struct__cef_before_download_callback_t
}

func (c CefBeforeDownloadCallbackT) Cont(downloadPath string, showDialog bool) {
	var downloadPathCString *C.char = C.CString(downloadPath)
	defer C.free(unsafe.Pointer(downloadPathCString))
	showDialogInt := 0
	if showDialog {
		showDialogInt = 1
	}
	C.callBeforeDownloadCallback_cont(c.Self, downloadPathCString, C.int(showDialogInt))
}

func (c CefBeforeDownloadCallbackT) Release() {
	C.releaseVoid(unsafe.Pointer(c.Self))
}

type CefDownloadItemCallbackT struct {
	Self *C.struct__cef_download_item_callback_t
}

func (c CefDownloadItemCallbackT) Cancel() {
	C.callDownloadItemCallback_cancel(c.Self)
}

func (c CefDownloadItemCallbackT) Release() {
	C.releaseVoid(unsafe.Pointer(c.Self))
}

type CefDownloadItemT struct {
	Self *C.struct__cef_download_item_t
}

func (c CefDownloadItemT) Release() {
	C.releaseVoid(unsafe.Pointer(c.Self))
}

func (c CefDownloadItemT) IsValid() bool {
	if C.callCefDownloadItem_is_valid(c.Self) == 0 {
		return false
	}
	return true
}
func (c CefDownloadItemT) IsInProgress() bool {
	if C.callCefDownloadItem_is_in_progress(c.Self) == 0 {
		return false
	}
	return true
}
func (c CefDownloadItemT) IsComplete() bool {
	if C.callCefDownloadItem_is_complete(c.Self) == 0 {
		return false
	}
	return true
}
func (c CefDownloadItemT) IsCanceled() bool {
	if C.callCefDownloadItem_is_canceled(c.Self) == 0 {
		return false
	}
	return true
}
func (c CefDownloadItemT) GetCurrentSpeed() int64 {
	return int64(C.callCefDownloadItem_get_current_speed(c.Self))
}
func (c CefDownloadItemT) GetPercentComplete() int {
	return int(C.callCefDownloadItem_get_percent_complete(c.Self))
}
func (c CefDownloadItemT) GetTotalBytes() int64 {
	return int64(C.callCefDownloadItem_get_total_bytes(c.Self))
}
func (c CefDownloadItemT) GetReceivedBytes() int64 {
	return int64(C.callCefDownloadItem_get_received_bytes(c.Self))
}
func (c CefDownloadItemT) GetStartTime() CefTimeT {
	return CefTimeT{C.callCefDownloadItem_get_start_time(c.Self)}
}
func (c CefDownloadItemT) GetEndTime() CefTimeT {
	return CefTimeT{C.callCefDownloadItem_get_end_time(c.Self)}
}
func (c CefDownloadItemT) GetFullPath() string {
	stringStruct := C.callCefDownloadItem_get_full_path(c.Self)
	if stringStruct == nil {
		return ""
	}
	defer C.cef_string_userfree_utf8_free(stringStruct)
	str := C.GoString(stringStruct.str)
	return str
}
func (c CefDownloadItemT) GetId() uint32 {
	return uint32(C.callCefDownloadItem_get_id(c.Self))
}
func (c CefDownloadItemT) GetUrl() string {
	stringStruct := C.callCefDownloadItem_get_url(c.Self)
	defer C.cef_string_userfree_utf8_free(stringStruct)
	str := C.GoString(stringStruct.str)
	return str
}
func (c CefDownloadItemT) GetSuggestedFileName() string {
	stringStruct := C.callCefDownloadItem_get_suggested_file_name(c.Self)
	defer C.cef_string_userfree_utf8_free(stringStruct)
	str := C.GoString(stringStruct.str)
	return str
}
func (c CefDownloadItemT) GetContentDisposition() string {
	stringStruct := C.callCefDownloadItem_get_content_disposition(c.Self)
	defer C.cef_string_userfree_utf8_free(stringStruct)
	str := C.GoString(stringStruct.str)
	return str
}
func (c CefDownloadItemT) GetMimeType() string {
	stringStruct := C.callCefDownloadItem_get_mime_type(c.Self)
	defer C.cef_string_userfree_utf8_free(stringStruct)
	str := C.GoString(stringStruct.str)
	return str
}

//export go_OnBeforeDownload
func go_OnBeforeDownload(
	self *C.struct__cef_download_handler_t,
	browser *C.struct__cef_browser_t,
	download_item *C.struct__cef_download_item_t,
	suggested_name *C.cef_string_utf8_t,
	callback *C.struct__cef_before_download_callback_t) {

	defer C.cef_string_userfree_utf8_free(suggested_name)
	str := C.GoString(suggested_name.str)

	if handler, ok := downloadHandlerMap[unsafe.Pointer(self)]; ok {
		handler.OnBeforeDownload(
			CefBrowserT{browser},
			CefDownloadItemT{download_item},
			str,
			CefBeforeDownloadCallbackT{callback},
		)
		return
	}
	CefBrowserT{browser}.Release()
	CefDownloadItemT{download_item}.Release()
	CefBeforeDownloadCallbackT{callback}.Release()

}

//export go_OnDownloadUpdated
func go_OnDownloadUpdated(
	self *C.struct__cef_download_handler_t,
	browser *C.struct__cef_browser_t,
	download_item *C.struct__cef_download_item_t,
	callback *C.struct__cef_download_item_callback_t) {

	if handler, ok := downloadHandlerMap[unsafe.Pointer(self)]; ok {
		handler.OnDownloadUpdated(
			CefBrowserT{browser},
			CefDownloadItemT{download_item},
			CefDownloadItemCallbackT{callback},
		)
		return
	}
	CefBrowserT{browser}.Release()
	CefDownloadItemT{download_item}.Release()
	CefDownloadItemCallbackT{callback}.Release()

}

func NewDownloadHandlerT(download DownloadHandler) DownloadHandlerT {
	var handler DownloadHandlerT
	handler.CStruct = (*C.struct__cef_download_handler_t)(
		C.calloc(1, C.sizeof_struct__cef_download_handler_t))
	C.initialize_download_handler(handler.CStruct)
	go_AddRef(unsafe.Pointer(handler.CStruct))
	downloadHandlerMap[unsafe.Pointer(handler.CStruct)] = download
	return handler
}
