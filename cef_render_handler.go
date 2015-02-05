package chrome

/*
#include <stdlib.h>
#include "cef_base.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_render_handler_capi.h"
extern void initialize_render_handler(struct _cef_render_handler_t* renderHandler);
*/
import "C"
import (
	log "github.com/cihub/seelog"
	"unsafe"
)

var renderHandlerMap = make(map[unsafe.Pointer]RenderHandler)

type RenderHandler interface {
	GetRenderHandler() RenderHandlerT
}

type RenderHandlerT struct {
	CStruct *C.struct__cef_render_handler_t
}

func (r RenderHandlerT) AddRef() {
	AddRef(unsafe.Pointer(r.CStruct))
}
func (r RenderHandlerT) Release() {
	Release(unsafe.Pointer(r.CStruct))
}

func NewRenderHandlerT(render RenderHandler) RenderHandlerT {
	var handler RenderHandlerT
	handler.CStruct = (*C.struct__cef_render_handler_t)(
		C.calloc(1, C.sizeof_struct__cef_render_handler_t))
	log.Info("initialize RenderHandler")

	C.initialize_render_handler(handler.CStruct)
	go_AddRef(unsafe.Pointer(handler.CStruct))
	renderHandlerMap[unsafe.Pointer(handler.CStruct)] = render
	return handler
}
