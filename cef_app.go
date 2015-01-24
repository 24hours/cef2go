// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go
// Website: https://github.com/fromkeith/cef2go

package chrome

/*
#cgo CFLAGS: -I./
#include <stdlib.h>
#include "include/capi/cef_app_capi.h"
extern void initialize_app_handler(cef_app_t* app);
*/
import "C"
import (
	log "github.com/cihub/seelog"
	"unsafe"
)

var knownAppHandlers = make(map[unsafe.Pointer]AppHandler)

// create the underlying C structure for an App Handler.
func NewAppHandlerT(handler AppHandler) AppHandlerT {
	var a AppHandlerT
	a.CStruct = (*C.cef_app_t)(C.calloc(1, C.sizeof_cef_app_t))
	log.Info("Initialize App Handler")
	C.initialize_app_handler(a.CStruct)

	knownAppHandlers[unsafe.Pointer(a.CStruct)] = handler
	return a
}

//export go_GetBrowserProcessHandler
func go_GetBrowserProcessHandler(self *C.cef_app_t) *C.struct__cef_browser_process_handler_t {
	if handler, ok := knownAppHandlers[unsafe.Pointer(self)]; ok {
		bph := handler.GetBrowserProcessHandler()
		if bph != nil {
			return bph.GetBrowserProcessHandlerT().CStruct
		}
	}
	return nil
}

type AppHandler interface {
	// TODO implement these thing
	OnBeforeCommandLineProcessing(processType string, commandLine CommandLineT)
	OnRegisterCustomSchemes()
	GetResourceBundleHandler()
	GetRenderProcessHandler()
	// called to get the underlying c struct.
	GetAppHandlerT() AppHandlerT
	GetBrowserProcessHandler() BrowserProcessHandler
}
type AppHandlerT struct {
	CStruct *C.cef_app_t
}

//base Handler
type BaseAppHandler struct {
	handler               AppHandlerT
	browserProcessHandler BrowserProcessHandler
}

func (app *BaseAppHandler) haveBase() bool { return true }
func (app *BaseAppHandler) OnBeforeCommandLineProcessing(processType string,
	commandLine CommandLineT) {
}
func (app *BaseAppHandler) OnRegisterCustomSchemes()  {}
func (app *BaseAppHandler) GetResourceBundleHandler() {}
func (app *BaseAppHandler) GetBrowserProcessHandler() BrowserProcessHandler {
	app.browserProcessHandler.GetBrowserProcessHandlerT().AddRef()
	return app.browserProcessHandler
}
func (app *BaseAppHandler) GetRenderProcessHandler() {}
func (app *BaseAppHandler) GetAppHandlerT() AppHandlerT {
	return app.handler
}
