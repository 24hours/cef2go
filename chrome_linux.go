// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package chrome

/*
#cgo CFLAGS: -I./
#cgo pkg-config: --libs --cflags gtk+-2.0
#include <stdlib.h>
#include <string.h>
#include <gtk/gtk.h>
#include "include/capi/cef_app_capi.h"

struct gtkInfo{
	int width;
	int height;
};

struct gtkInfo getBound(GtkWidget *g){
	struct gtkInfo ret;
	gtk_window_get_default_size(GTK_WINDOW(g), (gint*)(&ret.width), (gint*)(&ret.height));

	return ret;
}
*/
import "C"
import "unsafe"

import (
	log "github.com/cihub/seelog"
	"os"
)

var _Argv []*C.char = make([]*C.char, len(os.Args))

type WindowInfo struct {
	Ptr                 unsafe.Pointer
	WindowlessRendering int
	Height              int
	Width               int
}

func FillMainArgs(mainArgs *C.struct__cef_main_args_t,
	appHandle unsafe.Pointer) {
	// On Linux appHandle is nil.
	log.Debug("FillMainArgs, argv=", os.Args)
	for i, arg := range os.Args {
		_Argv[C.int(i)] = C.CString(arg)
	}
	mainArgs.argc = C.int(len(os.Args))
	mainArgs.argv = &_Argv[0]
}

func FillWindowInfo(windowInfo *C.cef_window_info_t, hwnd WindowInfo) {
	if hwnd.Ptr != nil {
		var info C.struct_gtkInfo = C.getBound((*C.GtkWidget)(hwnd.Ptr))
		windowInfo.parent_window = C.ulong(hwnd.Hdl)

		windowInfo.x = 0
		windowInfo.y = 0
		windowInfo.width = C.uint(info.width)
		windowInfo.height = C.uint(info.height)
	}
	windowInfo.windowless_rendering_enabled = C.int(hwnd.WindowlessRendering)
	windowInfo.height = C.uint(hwnd.Height)
	windowInfo.width = C.uint(hwnd.Width)

}
