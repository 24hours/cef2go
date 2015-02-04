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
import (
	log "github.com/cihub/seelog"
	"unsafe"
)

var downloadHandlerMap = make(map[unsafe.Pointer]DownloadHandler)

type DownloadHandler interface {
	OnBeforeDownload(browser Browser, downloadItem DownloadItem, suggestedName string, callback BeforeDownloadCallback)
	OnDownloadUpdated(browser Browser, downloadItem DownloadItem, callback DownloadItemCallback)

	GetDownloadHandlerT() DownloadHandlerT
}

type DownloadHandlerT struct {
	CStruct *C.struct__cef_download_handler_t
}

//TODO: implement house keeping code elsewhere
func (r DownloadHandlerT) AddRef() {
	AddRef(unsafe.Pointer(r.CStruct))
}
func (r DownloadHandlerT) Release() {
	Release(unsafe.Pointer(r.CStruct))
}

type CefTimeT struct {
	Self C.cef_time_t
}

type BeforeDownloadCallback struct {
	Self *C.struct__cef_before_download_callback_t
}

func (c BeforeDownloadCallback) Cont(downloadPath string, showDialog bool) {
	var downloadPathCString *C.char = C.CString(downloadPath)
	defer C.free(unsafe.Pointer(downloadPathCString))
	showDialogInt := 0
	if showDialog {
		showDialogInt = 1
	}
	C.callBeforeDownloadCallback_cont(c.Self, downloadPathCString, C.int(showDialogInt))
}

func (c BeforeDownloadCallback) Release() {
	C.releaseVoid(unsafe.Pointer(c.Self))
}

type DownloadItemCallback struct {
	Self *C.struct__cef_download_item_callback_t
}

func (c DownloadItemCallback) Cancel() {
	C.callDownloadItemCallback_cancel(c.Self)
}

func (c DownloadItemCallback) Release() {
	C.releaseVoid(unsafe.Pointer(c.Self))
}

type DownloadItem struct {
	Self *C.struct__cef_download_item_t
}

func (c DownloadItem) Release() {
	C.releaseVoid(unsafe.Pointer(c.Self))
}

func (c DownloadItem) IsValid() bool {
	if C.callCefDownloadItem_is_valid(c.Self) == 0 {
		return false
	} else {
		return true
	}
}

func (c DownloadItem) IsInProgress() bool {
	if C.callCefDownloadItem_is_in_progress(c.Self) == 0 {
		return false
	} else {
		return true
	}
}
func (c DownloadItem) IsComplete() bool {
	if C.callCefDownloadItem_is_complete(c.Self) == 0 {
		return false
	} else {
		return true
	}
}

func (c DownloadItem) IsCanceled() bool {
	if C.callCefDownloadItem_is_canceled(c.Self) == 0 {
		return false
	} else {
		return true
	}
}

func (c DownloadItem) GetCurrentSpeed() int64 {
	return int64(C.callCefDownloadItem_get_current_speed(c.Self))
}
func (c DownloadItem) GetPercentComplete() int {
	return int(C.callCefDownloadItem_get_percent_complete(c.Self))
}
func (c DownloadItem) GetTotalBytes() int64 {
	return int64(C.callCefDownloadItem_get_total_bytes(c.Self))
}
func (c DownloadItem) GetReceivedBytes() int64 {
	return int64(C.callCefDownloadItem_get_received_bytes(c.Self))
}
func (c DownloadItem) GetStartTime() CefTimeT {
	return CefTimeT{C.callCefDownloadItem_get_start_time(c.Self)}
}
func (c DownloadItem) GetEndTime() CefTimeT {
	return CefTimeT{C.callCefDownloadItem_get_end_time(c.Self)}
}
func (c DownloadItem) GetFullPath() string {
	stringStruct := C.callCefDownloadItem_get_full_path(c.Self)
	if stringStruct == nil {
		return ""
	}
	defer C.cef_string_userfree_utf8_free(stringStruct)
	str := C.GoString(stringStruct.str)
	return str
}
func (c DownloadItem) GetId() uint32 {
	return uint32(C.callCefDownloadItem_get_id(c.Self))
}
func (c DownloadItem) GetUrl() string {
	stringStruct := C.callCefDownloadItem_get_url(c.Self)
	defer C.cef_string_userfree_utf8_free(stringStruct)
	str := C.GoString(stringStruct.str)
	return str
}
func (c DownloadItem) GetSuggestedFileName() string {
	stringStruct := C.callCefDownloadItem_get_suggested_file_name(c.Self)
	defer C.cef_string_userfree_utf8_free(stringStruct)
	str := C.GoString(stringStruct.str)
	return str
}
func (c DownloadItem) GetContentDisposition() string {
	stringStruct := C.callCefDownloadItem_get_content_disposition(c.Self)
	defer C.cef_string_userfree_utf8_free(stringStruct)
	str := C.GoString(stringStruct.str)
	return str
}
func (c DownloadItem) GetMimeType() string {
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

	defer Browser{browser}.Release()
	defer DownloadItem{download_item}.Release()
	defer BeforeDownloadCallback{callback}.Release()

	if handler, ok := downloadHandlerMap[unsafe.Pointer(self)]; ok {
		handler.OnBeforeDownload(
			Browser{browser},
			DownloadItem{download_item},
			str,
			BeforeDownloadCallback{callback},
		)
	}
	return

}

//export go_OnDownloadUpdated
func go_OnDownloadUpdated(
	self *C.struct__cef_download_handler_t,
	browser *C.struct__cef_browser_t,
	download_item *C.struct__cef_download_item_t,
	callback *C.struct__cef_download_item_callback_t) {

	defer Browser{browser}.Release()
	defer DownloadItem{download_item}.Release()
	defer DownloadItemCallback{callback}.Release()

	if handler, ok := downloadHandlerMap[unsafe.Pointer(self)]; ok {
		handler.OnDownloadUpdated(
			Browser{browser},
			DownloadItem{download_item},
			DownloadItemCallback{callback},
		)
	}
	return
}

func NewDownloadHandlerT(download DownloadHandler) DownloadHandlerT {
	var handler DownloadHandlerT
	handler.CStruct = (*C.struct__cef_download_handler_t)(
		C.calloc(1, C.sizeof_struct__cef_download_handler_t))
	log.Info("initialize LifeSpanHandler")

	C.initialize_download_handler(handler.CStruct)
	go_AddRef(unsafe.Pointer(handler.CStruct))
	downloadHandlerMap[unsafe.Pointer(handler.CStruct)] = download
	return handler
}
