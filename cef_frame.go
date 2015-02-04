// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cefcapi

package chrome

/*
#include <stdlib.h>
#include "cef_base.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_frame_capi.h"
extern int cef_frame_t_is_valid(struct _cef_frame_t* self);
extern void cef_frame_t_undo(struct _cef_frame_t* self);
extern void cef_frame_t_redo(struct _cef_frame_t* self);
extern void cef_frame_t_cut(struct _cef_frame_t* self);
extern void cef_frame_t_copy(struct _cef_frame_t* self);
extern void cef_frame_t_paste(struct _cef_frame_t* self);
extern void cef_frame_t_del(struct _cef_frame_t* self);
extern void cef_frame_t_select_all(struct _cef_frame_t* self);
extern void cef_frame_t_view_source(struct _cef_frame_t* self);
extern void cef_frame_t_get_source(struct _cef_frame_t* self,struct _cef_string_visitor_t* visitor);
extern void cef_frame_t_get_text(struct _cef_frame_t* self,struct _cef_string_visitor_t* visitor);
extern void cef_frame_t_load_request(struct _cef_frame_t* self,struct _cef_request_t* request);
extern void cef_frame_t_load_url(struct _cef_frame_t* self,const char* url);
extern void cef_frame_t_load_string(struct _cef_frame_t* self,const char* string_val, const char* url);
extern void cef_frame_t_execute_java_script(struct _cef_frame_t* self,const char* code, const char* script_url,int start_line);
extern int cef_frame_t_is_main(struct _cef_frame_t* self);
extern int cef_frame_t_is_focused(struct _cef_frame_t* self);
extern cef_string_utf8_t* cef_frame_t_get_name(struct _cef_frame_t* self);
extern int64 cef_frame_t_get_identifier(struct _cef_frame_t* self);
extern struct _cef_frame_t* cef_frame_t_get_parent(struct _cef_frame_t* self);
extern cef_string_utf8_t* cef_frame_t_get_url(struct _cef_frame_t* self);
extern struct _cef_browser_t* cef_frame_t_get_browser(struct _cef_frame_t* self);
extern struct _cef_v8context_t* cef_frame_t_get_v8context(struct _cef_frame_t* self);
extern void cef_frame_t_visit_dom(struct _cef_frame_t* self,struct _cef_domvisitor_t* visitor);
*/
import "C"

import (
	"unsafe"
)

type Frame struct {
	CStruct *C.struct__cef_frame_t
}

type CefStringVisitorT struct {
	CStruct *C.struct__cef_string_visitor_t
}

type CefV8ContextT struct {
	CStruct *C.struct__cef_v8context_t
}

type CefDomVisitorT struct {
	CStruct *C.struct__cef_domvisitor_t
}

func (b Frame) Release() {
	C.releaseVoid(unsafe.Pointer(b.CStruct))
}

func (b Frame) AddRef() {
	C.add_refVoid(unsafe.Pointer(b.CStruct))
}

func (f Frame) IsValid() bool {
	return C.cef_frame_t_is_valid(f.CStruct) == C.int(1)
}
func (f Frame) Undo() {
	C.cef_frame_t_undo(f.CStruct)
}
func (f Frame) Redo() {
	C.cef_frame_t_redo(f.CStruct)
}
func (f Frame) Cut() {
	C.cef_frame_t_cut(f.CStruct)
}
func (f Frame) Copy() {
	C.cef_frame_t_copy(f.CStruct)
}
func (f Frame) Paste() {
	C.cef_frame_t_paste(f.CStruct)
}
func (f Frame) Del() {
	C.cef_frame_t_del(f.CStruct)
}
func (f Frame) SelectAll() {
	C.cef_frame_t_select_all(f.CStruct)
}
func (f Frame) ViewSource() {
	C.cef_frame_t_view_source(f.CStruct)
}
func (f Frame) GetSource(visitor CefStringVisitorT) {
	C.cef_frame_t_get_source(f.CStruct, visitor.CStruct)
}
func (f Frame) GetText(visitor CefStringVisitorT) {
	C.cef_frame_t_get_text(f.CStruct, visitor.CStruct)
}
func (f Frame) LoadRequest(request CefRequestT) {
	C.cef_frame_t_load_request(f.CStruct, request.CStruct)
}
func (f Frame) LoadUrl(url string) {
	cString := C.CString(url)
	defer C.free(unsafe.Pointer(cString))

	C.cef_frame_t_load_url(f.CStruct, cString)
}
func (f Frame) LoadString(contents, url string) {
	contentsCString := C.CString(contents)
	defer C.free(unsafe.Pointer(contentsCString))
	urlCString := C.CString(url)
	defer C.free(unsafe.Pointer(urlCString))

	C.cef_frame_t_load_string(f.CStruct, contentsCString, urlCString)
}
func (f Frame) ExecuteJavaScript(code, url string, startLine int) {
	codeCString := C.CString(code)
	defer C.free(unsafe.Pointer(codeCString))
	urlCString := C.CString(url)
	defer C.free(unsafe.Pointer(urlCString))

	C.cef_frame_t_execute_java_script(f.CStruct, codeCString, urlCString, C.int(startLine))
}
func (f Frame) IsMain() bool {
	return C.cef_frame_t_is_main(f.CStruct) == C.int(1)
}
func (f Frame) IsFocused() bool {
	return C.cef_frame_t_is_focused(f.CStruct) == C.int(1)
}
func (f Frame) GetName() string {
	cefString := C.cef_frame_t_get_name(f.CStruct)
	defer C.cef_string_userfree_utf8_free(cefString)
	return C.GoString(cefString.str)
}
func (f Frame) GetIdentifier() int64 {
	return int64(C.cef_frame_t_get_identifier(f.CStruct))
}
func (f Frame) GetParent() Frame {
	return Frame{C.cef_frame_t_get_parent(f.CStruct)}
}
func (f Frame) GetUrl() string {
	cefString := C.cef_frame_t_get_url(f.CStruct)
	defer C.cef_string_userfree_utf8_free(cefString)
	return C.GoString(cefString.str)
}
func (f Frame) GetBrowser() Browser {
	return Browser{C.cef_frame_t_get_browser(f.CStruct)}
}
func (f Frame) GetV8context() CefV8ContextT {
	return CefV8ContextT{C.cef_frame_t_get_v8context(f.CStruct)}
}
func (f Frame) VisitDom(visitor CefDomVisitorT) {
	C.cef_frame_t_visit_dom(f.CStruct, visitor.CStruct)
}
