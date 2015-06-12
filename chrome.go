// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

//TODO : figure out why crash on shutdown()
// 	     Shutdown() already calld in UI thread,
// 			 but it complain regardless, maybe a bug in CEF
//			 leave it untill next release come out
//     : dragging will cause CEF to crash
package chrome

/*
CEF capi fixes
--------------
1. In cef_string.h:
    this => typedef cef_string_utf16_t cef_string_t;
    to => #define cef_string_t cef_string_utf16_t
2. In cef_export.h:
    #elif defined(COMPILER_GCC)
    #define CEF_EXPORT __attribute__ ((visibility("default")))
    #ifdef OS_WIN
    #define CEF_CALLBACK __stdcall
    #else
    #define CEF_CALLBACK
    #endif
*/

/*
#include <stdlib.h>
#include <string.h>
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_client_capi.h"
#include "cef_base.h"
*/
import "C"
import "unsafe"
import (
	"errors"
	log "github.com/cihub/seelog"
	"os"
)

var logger log.LoggerInterface
var execute_called bool
var _MainArgs *C.struct__cef_main_args_t

// Sandbox is disabled. Including the "cef_sandbox.lib"
// library results in lots of GCC warnings/errors. It is
// compatible only with VS 2010. It would be required to
// build it using GCC. Add -lcef_sandbox to LDFLAGS.
// capi doesn't expose sandbox functions, you need do add
// these before import "C":
// void* cef_sandbox_info_create();
// void cef_sandbox_info_destroy(void* sandbox_info);
var _SandboxInfo unsafe.Pointer

func init() {
	DisableLog()
}

func NewWindowInfo(height, width int) WindowInfo {
	ret := WindowInfo{}
	ret.Height = height
	ret.Width = width

	return ret
}

// TODO : Add this line back
//func ExecuteProcess(programHandle unsafe.Pointer, appHandler AppHandler) int {
func ExecuteProcess(programHandle unsafe.Pointer, appHandler interface{}) int {
	var appHandlerT *C.cef_app_t
	if appHandler == nil {
		appHandlerT = nil
	} else {
		// appHandlerT = appHandler.GetAppHandlerT().CStruct
		// go_AddRef(unsafe.Pointer(appHandlerT))
	}

	_MainArgs = (*C.struct__cef_main_args_t)(C.calloc(1, C.sizeof_struct__cef_main_args_t))
	FillMainArgs(_MainArgs, programHandle)
	// Sandbox info needs to be passed to both cef_execute_process()
	// and cef_initialize().
	// OFF: _SandboxInfo = C.cef_sandbox_info_create()

	var exitCode C.int = C.cef_execute_process(_MainArgs, appHandlerT, _SandboxInfo)
	if exitCode >= 0 {
		os.Exit(int(exitCode))
	}
	execute_called = true
	return int(exitCode)
}

// TODO : add this back
// func Initialize(settings Settings, appHandler AppHandler) int {
func Initialize(settings Settings, appHandler interface{}) int {
	if execute_called == false {
		// If cef_initialize called before  cef_execute_process
		// then it would result in creation of infinite number of
		// processes. See Issue 1199 in CEF:
		// https://code.google.com/p/chromiumembedded/issues/detail?id=1199
		log.Critical("missing call to ExecuteProcess")
		panic("Missing Call to ExecuteProcess")
	}

	var appHandlerT *C.cef_app_t
	if appHandler == nil {
		appHandlerT = nil
	} else {
		// appHandlerT = appHandler.GetAppHandlerT().CStruct
	}
	ret := C.cef_initialize(_MainArgs, settings.toC(), appHandlerT, _SandboxInfo)
	return int(ret)
}

// TODO : add this back
// func CreateBrowserSync(hwnd WindowInfo,
// 	clientHandler ClientHandler,
// 	browserSettings BrowserSettings,
// 	url string) (*Browser, error) {

func CreateBrowserSync(hwnd WindowInfo,
	clientHandler interface{},
	browserSettings BrowserSettings,
	url string) (*interface{}, error) {

	return createBrowser(hwnd, clientHandler, browserSettings, url, true)
}

// TODO : add this back
// func CreateBrowserAsync(hwnd WindowInfo,
// 	clientHandler ClientHandler,
// 	browserSettings BrowserSettings,
// 	url string) error {
func CreateBrowserAsync(hwnd WindowInfo,
	clientHandler interface{},
	browserSettings BrowserSettings,
	url string) error {

	_, err := createBrowser(hwnd, clientHandler, browserSettings, url, false)
	return err
}

// TODO : add this back
// func createBrowser(hwnd WindowInfo,
// 	clientHandler ClientHandler,
// 	browserSettings BrowserSettings,
// 	url string,
// async bool) (*Browser, error) {

func createBrowser(hwnd WindowInfo,
	clientHandler interface{},
	browserSettings BrowserSettings,
	url string,
	async bool) (*interface{}, error) {

	// Initialize cef_window_info_t structure.
	var windowInfo *C.cef_window_info_t
	windowInfo = (*C.cef_window_info_t)(C.calloc(1, C.sizeof_cef_window_info_t))
	FillWindowInfo(windowInfo, hwnd)

	// url
	var cefUrl *C.cef_string_t
	cefUrl = (*C.cef_string_t)(C.calloc(1, C.sizeof_cef_string_t))
	toCefStringCopy(url, cefUrl)

	var ch *C.struct__cef_client_t
	// if clientHandler == nil {
	// 	ch = nil
	// 	if hwnd.WindowlessRendering == 1 {
	// 		log.Critical("RenderHandler must be implemented for Windowless Rendering to work")
	// 		return nil, errors.New("RenderHandler must be implemented for Windowless Rendering to work")
	// 	}
	// } else {
	// 	//registering and activating multiple handler
	// 	if _, ok := clientHandler.(ClientHandler); ok {
	// 		log.Debug("Registering Client hanlder")
	// 		clientHandler.SetClientHandlerT(NewClientHandlerT(clientHandler))
	// 	} else {
	// 		log.Critical("clientHandler must implement interface ClientHandler")
	// 		return nil, errors.New("ClientHandler not implemented")
	// 	}

	// 	if lsh, ok := clientHandler.(LifeSpanHandler); ok {
	// 		log.Debug("Registering Life Span handler ")
	// 		clientHandler.SetLifeSpanHandler(NewLifeSpanHandlerT(lsh))
	// 	}

	// 	if rqh, ok := clientHandler.(RequestHandler); ok {
	// 		log.Debug("Registering Request handler ")
	// 		clientHandler.SetRequestHandler(NewRequestHandlerT(rqh))
	// 	}

	// 	if dsp, ok := clientHandler.(DisplayHandler); ok {
	// 		log.Debug("Registering Display handler ")
	// 		clientHandler.SetDisplayHandler(NewDisplayHandlerT(dsp))
	// 	}

	// 	if dl, ok := clientHandler.(DownloadHandler); ok {
	// 		log.Debug("Registering Download handler ")
	// 		clientHandler.SetDownloadHandler(NewDownloadHandlerT(dl))
	// 	}

	// 	if rn, ok := clientHandler.(RenderHandler); ok {
	// 		log.Debug("Registering Render handler ")
	// 		clientHandler.SetRenderHandler(NewRenderHandlerT(rn))
	// 	} else {
	// 		if hwnd.WindowlessRendering == 1 {
	// 			log.Critical("RenderHandler must be implemented for Windowless Rendering to work")
	// 			return nil, errors.New("RenderHandler must be implemented for Windowless Rendering to work")
	// 		}
	// 	}

	// 	ch = clientHandler.GetClientHandlerT().CStruct
	// 	go_AddRef(unsafe.Pointer(ch))
	// }

	if async == false {
		result := C.cef_browser_host_create_browser(
			windowInfo,
			ch,
			cefUrl,
			browserSettings.toC(),
			nil,
		)

		if result != C.int(1) {
			log.Error("C.cef_browser_host_create_browser return :", result)
			return nil, errors.New("Unknown failure in CEF")
		} else {
			return nil, nil
		}
	} else {
		// Do not create the browser synchronously using the
		// cef_browser_host_create_browser_sync() function, as
		// it is unreliable. Instead obtain browser object in
		// life_span_handler::on_after_created. In that callback
		// keep CEF browser objects in a global map (cef window
		// handle -> cef browser) and introduce
		// a GetBrowserByWindowHandle() function. This function
		// will first guess the CEF window handle using for example
		// WinAPI functions and then search the global map of cef
		// browser objects.

		// browser := &Browser{}
		// browser.CStruct = C.cef_browser_host_create_browser_sync(
		// 	windowInfo,
		// 	ch,
		// 	cefUrl,
		// 	browserSettings.toC(),
		// 	nil,
		// )
		// if browser.CStruct != nil {
		// 	log.Error("C.cef_browser_host_create_browser_sync fail to return browser pointer")
		// 	return nil, errors.New("Unknown failure in CEF")
		// } else {
		// 	return browser, nil
		// }
		return nil, errors.New("create browser sync will fail no matter what")
	}
}

func RunMessageLoop() {
	log.Info("RunMessageLoop")
	C.cef_run_message_loop()
}

func QuitMessageLoop() {
	log.Info("QuitMessageLoop")
	C.cef_quit_message_loop()
}

func Shutdown() {
	log.Info("Shutdown")
	C.free(unsafe.Pointer(_MainArgs))
	C.cef_shutdown()
	// OFF: cef_sandbox_info_destroy(_SandboxInfo)
}
