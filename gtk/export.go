// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package gtk

//#include <gtk/gtk.h>
import "C"
import "unsafe"
import log "github.com/cihub/seelog"

//export _GoDestroySignal
func _GoDestroySignal(widget *C.GtkWidget, data C.gpointer) {
	ptr := uintptr(unsafe.Pointer(widget))
	if callback, ok := destroySignalCallbacks[ptr]; ok {
		delete(destroySignalCallbacks, ptr)
		callback()
	} else {
		log.Warn("_GoDestroySignal failed, callback not found")
	}
}
