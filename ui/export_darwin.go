// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package ui

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include <Cocoa/Cocoa.h>
*/
import "C"
import "unsafe"
import log "github.com/cihub/seelog"

//export _GoDestroySignal
func _GoDestroySignal(window unsafe.Pointer) {
	ptr := uintptr(window)
	if callback, ok := destroySignalCallbacks[ptr]; ok {
		delete(destroySignalCallbacks, ptr)
		callback()
	} else {
		log.Warn(" _GoDestroySignal failed, callback not found")
	}

}
