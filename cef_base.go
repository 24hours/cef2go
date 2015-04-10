// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go
// Website: https://github.com/fromkeith/cef2go

package chrome

/*
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_app_capi.h"
#include "cef_base.h"

extern int releaseVoid(void * self);
*/
import "C"
import (
	"errors"
	"fmt"
	log "github.com/cihub/seelog"
	"sync"
	"unsafe"
)

// a niave memory management.
// allows us to keep track of allocated resources, and their counts
// furthermore, the deconstructor lets us free any go resources once
// their C versions have been released
// General info about the ref count in CEF:
//      http://code.google.com/p/chromiumembedded/wiki/UsingTheCAPI
type MemoryManagedBridge struct {
	Count         int
	Deconstructor func(it unsafe.Pointer)
	Name          string
}

var (
	memoryBridge = make(map[unsafe.Pointer]MemoryManagedBridge)
	refCountLock sync.Mutex
)

//export go_AddRef
func go_AddRef(it unsafe.Pointer) {
	if it == nil {
		return
	}
	refCountLock.Lock()
	defer refCountLock.Unlock()

	if m, ok := memoryBridge[it]; ok {
		m.Count++
		memoryBridge[it] = m
		// log.Debug("Adding new reference: ", it)
		return
	}
}

//export go_Release
func go_Release(it unsafe.Pointer) int {
	if it == nil {
		return 0
	}
	refCountLock.Lock()
	defer refCountLock.Unlock()

	if m, ok := memoryBridge[it]; ok {
		m.Count--
		if m.Count == 0 {
			if m.Deconstructor != nil {
				m.Deconstructor(it)
			}
			C.free(it)
			delete(memoryBridge, it)
			return 1
		} else {
			//log.Debug("Reduce reference count for: ", it)
			memoryBridge[it] = m
		}
		return 0
	}
	return 0
}

//export go_HasOneReferenceCount
func go_HasOneReferenceCount(it unsafe.Pointer) int {
	refCountLock.Lock()
	defer refCountLock.Unlock()
	if m, ok := memoryBridge[it]; ok {
		if m.Count == 1 {
			return 1
		}
		return 0
	}
	return 0
}

func Release(w unsafe.Pointer) int {
	return int(C.releaseVoid(w))
}

func AddRef(it unsafe.Pointer) {
	C.add_refVoid(it)
}

//export go_CreateRef
func go_CreateRef(it unsafe.Pointer, name *C.char) {
	goname := ""
	if name != nil {
		goname = C.GoString(name)
	}
	log.Info("Create Reference for ", goname)
	CreateRef(it, goname)
}

func CreateRef(it unsafe.Pointer, name string) {
	refCountLock.Lock()
	defer refCountLock.Unlock()

	if _, ok := memoryBridge[it]; !ok {
		var m MemoryManagedBridge
		m.Deconstructor = nil
		m.Name = name
		memoryBridge[it] = m
		return
	}
}

// TODO : This seems to be useless
func RegisterDestructor(it unsafe.Pointer, decon func(it unsafe.Pointer)) bool {
	refCountLock.Lock()
	defer refCountLock.Unlock()

	if m, ok := memoryBridge[it]; ok {
		m.Deconstructor = decon
		memoryBridge[it] = m
		return true
	}
	return false
}

func DumpRefs() {
	fmt.Println("Dumping reference : ")
	refCountLock.Lock()
	defer refCountLock.Unlock()

	for k, v := range memoryBridge {
		fmt.Println("%X : %#v", k, v)
	}
}

func clearAllRefs() {
	refCountLock.Lock()
	defer refCountLock.Unlock()

	for k, v := range memoryBridge {
		v.Count = 1
		go_Release(k)
	}
}

func getRef(it unsafe.Pointer) (error, MemoryManagedBridge) {
	if m, ok := memoryBridge[it]; ok {
		return nil, m
	} else {
		mem := MemoryManagedBridge{}
		return errors.New("No reference found"), mem
	}
}
