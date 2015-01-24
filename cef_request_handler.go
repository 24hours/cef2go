// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cef2go

package chrome

/*
#include <stdlib.h>
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_request_handler_capi.h"
extern void initialize_request_handler(struct _cef_request_handler_t * requestHandler);
extern void cef_allow_certificate_error_callback_t_cont(struct _cef_allow_certificate_error_callback_t* self, int allow);
*/
import "C"
import "unsafe"
import log "github.com/cihub/seelog"

var requestHandlerMap = make(map[unsafe.Pointer]RequestHandler)

// Wraps the callbacks done to _cef_request_handler_t (partial implementation of callbacks)
type RequestHandler interface {
	/** Triggered before browsing to page. Return 1 to cancel transition. 0 to continue. */
	OnBeforeBrowse(browser CefBrowserT, frame CefFrameT, request CefRequestT, isRedirect int) int
	OnBeforeResourceLoad(browser CefBrowserT, frame CefFrameT, request CefRequestT) int
	/** Triggered when we encounter an invalid SSL cert. Return 1 to and call callback.cont() to continue or cancel the request.
	  Return 0 to immediatly cancel the request.*/
	OnCertificateError(errorCode CefErrorCode, requestUrl string, errorCallback CefCertErrorCallbackT) int

	GetRequestHandler() RequestHandlerT
}

type RequestHandlerT struct {
	CStruct *C.struct__cef_request_handler_t
}

func (r RequestHandlerT) AddRef() {
	AddRef(unsafe.Pointer(r.CStruct))
}
func (r RequestHandlerT) Release() {
	Release(unsafe.Pointer(r.CStruct))
}

type CefCertErrorCallbackT struct {
	CStruct *C.struct__cef_allow_certificate_error_callback_t
}

func (c CefCertErrorCallbackT) Release() {
	Release(unsafe.Pointer(c.CStruct))
}

func (c CefCertErrorCallbackT) Cont(allow int) {
	C.cef_allow_certificate_error_callback_t_cont(c.CStruct, C.int(allow))
}

//export go_OnBeforeBrowse
func go_OnBeforeBrowse(
	self *C.struct__cef_request_handler_t,
	browser *C.struct__cef_browser_t,
	frame *C.struct__cef_frame_t,
	request *C.struct__cef_request_t,
	is_redirect int) int {

	defer CefBrowserT{browser}.Release()
	defer CefRequestT{request}.Release()
	defer CefFrameT{frame}.Release()
	if handler, ok := requestHandlerMap[unsafe.Pointer(self)]; ok {

		return handler.OnBeforeBrowse(
			CefBrowserT{browser},
			CefFrameT{frame},
			CefRequestT{request},
			is_redirect,
		)
	}

	return 0
}

//export go_OnBeforeResourceLoad
func go_OnBeforeResourceLoad(
	self *C.struct__cef_request_handler_t,
	browser *C.struct__cef_browser_t,
	frame *C.struct__cef_frame_t,
	request *C.struct__cef_request_t) int {
	defer CefBrowserT{browser}.Release()
	defer CefFrameT{frame}.Release()
	defer CefRequestT{request}.Release()
	if handler, ok := requestHandlerMap[unsafe.Pointer(self)]; ok {
		return handler.OnBeforeResourceLoad(
			CefBrowserT{browser},
			CefFrameT{frame},
			CefRequestT{request},
		)
	}

	return 0
}

//export go_GetResourceHandler
func go_GetResourceHandler(
	self *C.struct__cef_request_handler_t,
	browser *C.struct__cef_browser_t,
	frame *C.struct__cef_frame_t,
	request *C.struct__cef_request_t) *C.struct__cef_resource_handler_t {

	CefBrowserT{browser}.Release()
	CefFrameT{frame}.Release()
	CefRequestT{request}.Release()

	return nil
}

//export go_OnResourceRedirect
func go_OnResourceRedirect(
	self *C.struct__cef_request_handler_t,
	browser *C.struct__cef_browser_t,
	frame *C.struct__cef_frame_t,
	old_url *C.char,
	new_url *C.char) {

	CefBrowserT{browser}.Release()
	CefFrameT{frame}.Release()
}

//export go_GetAuthCredentials
func go_GetAuthCredentials(
	self *C.struct__cef_request_handler_t,
	browser *C.struct__cef_browser_t,
	frame *C.struct__cef_frame_t,
	isProxy int,
	host *C.char,
	port int,
	realm *C.char,
	scheme *C.char,
	callback *C.struct__cef_auth_callback_t) int {

	CefBrowserT{browser}.Release()
	CefFrameT{frame}.Release()
	Release(unsafe.Pointer(callback))
	return 1
}

//export go_OnQuotaRequest
func go_OnQuotaRequest(
	self *C.struct__cef_request_handler_t,
	browser *C.struct__cef_browser_t,
	origin_url *C.char,
	new_size int64,
	callback *C.struct__cef_quota_callback_t) int {
	CefBrowserT{browser}.Release()
	Release(unsafe.Pointer(callback))
	return 1
}

//export go_OnProtocolExecution
func go_OnProtocolExecution(
	self *C.struct__cef_request_handler_t,
	browser *C.struct__cef_browser_t,
	url *C.char,
	allow_os_execution unsafe.Pointer) {

	CefBrowserT{browser}.Release()
}

//export go_OnCertificateError
func go_OnCertificateError(
	self *C.struct__cef_request_handler_t,
	cert_error int, //C.enum_cef_errorcode_t,
	request_url *C.char,
	callback *C.struct__cef_allow_certificate_error_callback_t) int {
	defer CefCertErrorCallbackT{callback}.Release()
	if handler, ok := requestHandlerMap[unsafe.Pointer(self)]; ok {
		return handler.OnCertificateError(
			CefErrorCode(cert_error),
			C.GoString(request_url),
			CefCertErrorCallbackT{callback},
		)
	}

	return 0
}

//export go_OnBeforePluginLoad
func go_OnBeforePluginLoad(
	self *C.struct__cef_request_handler_t,
	browser *C.struct__cef_browser_t,
	url *C.char,
	policy_url *C.char,
	info *C.struct__cef_web_plugin_info_t) int {

	CefBrowserT{browser}.Release()
	Release(unsafe.Pointer(info))
	return 1
}

//export go_OnPluginCrashed
func go_OnPluginCrashed(
	self *C.struct__cef_request_handler_t,
	browser *C.struct__cef_browser_t,
	plugin_path *C.char) {
	CefBrowserT{browser}.Release()
}

//export go_OnRenderProcessTerminated
func go_OnRenderProcessTerminated(
	self *C.struct__cef_request_handler_t,
	browser *C.struct__cef_browser_t,
	status int, //C.enum_cef_termination_status_t
) {
	CefBrowserT{browser}.Release()
}

func NewRequestHandlerT(request RequestHandler) RequestHandlerT {
	var handler RequestHandlerT
	handler.CStruct = (*C.struct__cef_request_handler_t)(
		C.calloc(1, C.sizeof_struct__cef_request_handler_t))

	log.Info("Initialize Request Handler")
	C.initialize_request_handler(handler.CStruct)

	go_AddRef(unsafe.Pointer(handler.CStruct))
	requestHandlerMap[unsafe.Pointer(handler.CStruct)] = request
	return handler
}
