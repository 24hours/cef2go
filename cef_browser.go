// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cefcapi

package chrome

/*
#include <stdlib.h>
#include "cef_base.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_browser_capi.h"
extern struct _cef_browser_host_t* cef_browser_t_get_host(struct _cef_browser_t* self);
extern int cef_browser_t_can_go_back(struct _cef_browser_t* self);
extern void cef_browser_t_go_back(struct _cef_browser_t* self);
extern int cef_browser_t_can_go_forward(struct _cef_browser_t* self);
extern void cef_browser_t_go_forward(struct _cef_browser_t* self);
extern int cef_browser_t_is_loading(struct _cef_browser_t* self);
extern void cef_browser_t_reload(struct _cef_browser_t* self);
extern void cef_browser_t_reload_ignore_cache(struct _cef_browser_t* self);
extern void cef_browser_t_stop_load(struct _cef_browser_t* self);
extern int cef_browser_t_get_identifier(struct _cef_browser_t* self);
extern int cef_browser_t_is_same(struct _cef_browser_t* self, struct _cef_browser_t* that);
extern int cef_browser_t_is_popup(struct _cef_browser_t* self);
extern int cef_browser_t_has_document(struct _cef_browser_t* self);
extern struct _cef_frame_t* cef_browser_t_get_main_frame(struct _cef_browser_t* self);
extern struct _cef_frame_t* cef_browser_t_get_focused_frame(struct _cef_browser_t* self);
extern struct _cef_frame_t* cef_browser_t_get_frame_byident(struct _cef_browser_t* self, int64 identifier);
extern struct _cef_frame_t* cef_browser_t_get_frame(struct _cef_browser_t* self, char * nameChar);
extern size_t cef_browser_t_get_frame_count(struct _cef_browser_t* self);
extern size_t cef_browser_t_get_frame_identifiers(struct _cef_browser_t* self, size_t count, int64 * ids);
extern void cef_browser_t_get_frame_names(struct _cef_browser_t* self, cef_string_list_t names);
extern int64 int64_array_get(int64* aa, size_t i);
*/
import "C"

import (
	"unsafe"
)

type Browser struct {
	CStruct *C.struct__cef_browser_t
}

func (b Browser) Release() {
	C.releaseVoid(unsafe.Pointer(b.CStruct))
}

func (b Browser) AddRef() {
	C.add_refVoid(unsafe.Pointer(b.CStruct))
}

func (b Browser) GetHost() CefBrowserHostT {
	hostStruct := C.cef_browser_t_get_host(b.CStruct)
	return CefBrowserHostT{hostStruct}
}

func (b Browser) CanGoBack() bool {
	return C.cef_browser_t_can_go_back(b.CStruct) == 1
}

func (b Browser) GoBack() {
	C.cef_browser_t_go_back(b.CStruct)
}

func (b Browser) CanGoForward() bool {
	return C.cef_browser_t_can_go_forward(b.CStruct) == 1
}

func (b Browser) GoForward() {
	C.cef_browser_t_go_forward(b.CStruct)
}

func (b Browser) IsLoading() bool {
	return C.cef_browser_t_is_loading(b.CStruct) == 1
}

func (b Browser) Reload() {
	C.cef_browser_t_reload(b.CStruct)
}

func (b Browser) ReloadIgnoreCache() {
	C.cef_browser_t_reload_ignore_cache(b.CStruct)
}

func (b Browser) StopLoad() {
	C.cef_browser_t_stop_load(b.CStruct)
}

func (b Browser) GetIdentifier() int {
	return int(C.cef_browser_t_get_identifier(b.CStruct))
}

func (b Browser) IsSame(other Browser) bool {
	return C.cef_browser_t_is_same(b.CStruct, other.CStruct) == 1
}

func (b Browser) IsPopup() bool {
	return C.cef_browser_t_is_popup(b.CStruct) == 1
}

func (b Browser) HasDocument() bool {
	return C.cef_browser_t_has_document(b.CStruct) == 1
}

func (b Browser) GetMainFrame() CefFrameT {
	return CefFrameT{C.cef_browser_t_get_main_frame(b.CStruct)}
}

func (b Browser) GetFocusedFrame() CefFrameT {
	return CefFrameT{C.cef_browser_t_get_focused_frame(b.CStruct)}
}

func (b Browser) GetFrameByIdent(identifier int64) CefFrameT {
	return CefFrameT{C.cef_browser_t_get_frame_byident(b.CStruct, C.int64(identifier))}
}

func (b Browser) GetFrame(name string) CefFrameT {
	cString := C.CString(name)
	defer C.free(unsafe.Pointer(cString))
	result := CefFrameT{C.cef_browser_t_get_frame(b.CStruct, cString)}
	return result
}

func (b Browser) GetFrameCount() int64 {
	return int64(C.cef_browser_t_get_frame_count(b.CStruct))
}

func (b Browser) GetFrameIdentifiers() []int64 {
	var count C.size_t = C.size_t(b.GetFrameCount())
	var ids *C.int64 = (*C.int64)(C.calloc(count, C.sizeof_int64))
	count = C.cef_browser_t_get_frame_identifiers(b.CStruct, count, ids)
	rCount := int64(count)
	result := make([]int64, rCount)
	var i int64
	for i = 0; i < rCount; i++ {
		result[i] = int64(C.int64_array_get(ids, C.size_t(i)))
	}
	C.free(unsafe.Pointer(ids))
	return result
}

func (b Browser) GetFrameNames() []string {
	nameList := C.cef_string_list_alloc()
	C.cef_browser_t_get_frame_names(b.CStruct, nameList)
	length := int(C.cef_string_list_size(nameList))
	goList := make([]string, length)
	for i := range goList {
		var cefName *C.cef_string_utf16_t = C.cef_string_userfree_utf16_alloc()
		didRet := C.cef_string_list_value(nameList, C.int(i), C.cefString16CastToCefString(cefName))
		if didRet == C.int(1) {
			nameUtf8 := C.cefStringToUtf8(C.cefString16CastToCefString(cefName))
			goList[i] = C.GoString(nameUtf8.str)
			C.cef_string_userfree_utf8_free(nameUtf8)
		}
		C.cef_string_userfree_utf16_free(cefName)
	}
	return goList
}

//func (b Browser) SendProcessMessage(targetProcess int, message *C.struct__cef_process_message_T) int
