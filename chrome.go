// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

//TODO : figure out why crash on shutdown()
// 	     Shutdown() already calld in UI thread,
// 		 but it complain regardless, maybe a bug in CEF
//		 leave it untill next release come out
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
#include "cef_app.h"
#include "cef_client.h"
*/
import "C"
import "unsafe"
import (
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

type WindowInfo struct {
	Ptr unsafe.Pointer
	Hdl uint64
}

func init() {
	logger, _ = log.LoggerFromWriterWithMinLevelAndFormat(os.Stdout, 0, "[%Level] %File:%Line: %Msg %n")
	log.ReplaceLogger(logger)
}

func ExecuteProcess(programHandle unsafe.Pointer, appHandler AppHandler) int {
	var appHandlerT *C.cef_app_t
	if appHandler == nil {
		appHandlerT = nil
	} else {
		appHandlerT = appHandler.GetAppHandlerT().CStruct
		go_AddRef(unsafe.Pointer(appHandlerT))
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

func Initialize(settings Settings, appHandler AppHandler) int {
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
		appHandlerT = appHandler.GetAppHandlerT().CStruct
		go_AddRef(unsafe.Pointer(appHandlerT))
	}
	ret := C.cef_initialize(_MainArgs, settings.toC(), appHandlerT, _SandboxInfo)
	return int(ret)
}

func CreateBrowser(hwnd WindowInfo,
	clientHandler ClientHandler,
	browserSettings BrowserSettings,
	url string) bool {

	// Initialize cef_window_info_t structure.
	var windowInfo *C.cef_window_info_t
	windowInfo = (*C.cef_window_info_t)(C.calloc(1, C.sizeof_cef_window_info_t))
	FillWindowInfo(windowInfo, hwnd)

	// url
	var cefUrl *C.cef_string_t
	cefUrl = (*C.cef_string_t)(C.calloc(1, C.sizeof_cef_string_t))
	toCefStringCopy(url, cefUrl)

	var ch *C.struct__cef_client_t
	if clientHandler == nil {
		ch = nil
	} else {
		//registering and activating multiple handler
		if _, ok := clientHandler.(ClientHandler); ok {
			log.Debug("Registering Client hanlder")
			clientHandler.SetClientHandlerT(NewClientHandlerT(clientHandler))
		} else {
			log.Critical("clientHandler must implement interface ClientHandler")
			panic("ClientHandler not implemented")
		}

		if lsh, ok := clientHandler.(LifeSpanHandler); ok {
			log.Debug("Registering Life Span handler ")
			clientHandler.SetLifeSpanHandler(NewLifeSpanHandlerT(lsh))
		}

		if rqh, ok := clientHandler.(RequestHandler); ok {
			log.Debug("Registering Request handler ")
			clientHandler.SetRequestHandler(NewRequestHandlerT(rqh))
		}

		ch = clientHandler.GetClientHandlerT().CStruct
		go_AddRef(unsafe.Pointer(ch))
	}

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

	result := C.cef_browser_host_create_browser(
		windowInfo,
		ch,
		cefUrl,
		browserSettings.toC(),
		nil,
	)

	return result == C.int(1)
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

func extractCefMultiMap(cefMapPointer C.cef_string_multimap_t) map[string][]string {
	numKeys := C.cef_string_multimap_size(cefMapPointer)
	goMap := make(map[string][]string)
	for i := 0; i < int(numKeys); i++ {
		var key *C.cef_string_utf16_t = C.cef_string_userfree_utf16_alloc()
		C.cef_string_multimap_key(cefMapPointer, C.int(i), C.cefString16CastToCefString(key))
		charKeyUtf8 := C.cefStringToUtf8(C.cefString16CastToCefString(key))
		goKey := C.GoString(charKeyUtf8.str)
		if _, ok := goMap[goKey]; ok {
			continue
		}
		numValsForKey := C.cef_string_multimap_find_count(cefMapPointer, C.cefString16CastToCefString(key))

		if numValsForKey >= 0 {
			goVals := make([]string, numValsForKey)
			for k := 0; k < int(numValsForKey); k++ {
				var val *C.cef_string_utf16_t = C.cef_string_userfree_utf16_alloc()
				C.cef_string_multimap_enumerate(cefMapPointer,
					C.cefString16CastToCefString(key), C.int(k), C.cefString16CastToCefString(val))
				charValUtf8 := C.cefStringToUtf8(C.cefString16CastToCefString(val))
				goVals[k] = C.GoString(charValUtf8.str)
				C.cef_string_userfree_utf8_free(charValUtf8)
				C.cef_string_userfree_utf16_free(val)
			}
			goMap[goKey] = goVals
		}
		C.cef_string_userfree_utf8_free(charKeyUtf8)
		C.cef_string_userfree_utf16_free(key)
	}
	return goMap
}
