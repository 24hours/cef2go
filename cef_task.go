// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cef2go

package chrome

/*
#include <stdlib.h>
#include "include/capi/cef_task_capi.h"
#include "include/internal/cef_types.h"
extern void initialize_task(struct _cef_task_t* self);
extern void helper_cef_post_task(cef_thread_id_t threadId, struct _cef_task_t* what);
*/
import "C"
import (
	log "github.com/cihub/seelog"
	"sync"
	"unsafe"
)

type ThreadId int

func OnThreadID(id ThreadId) bool {
	switch id {
	case TID_UI:
		return 1 == C.cef_currently_on(C.TID_UI)
	case TID_DB:
		return 1 == C.cef_currently_on(C.TID_DB)
	case TID_FILE:
		return 1 == C.cef_currently_on(C.TID_FILE)
	case TID_FILE_USER_BLOCKING:
		return 1 == C.cef_currently_on(C.TID_FILE_USER_BLOCKING)
	case TID_PROCESS_LAUNCHER:
		return 1 == C.cef_currently_on(C.TID_PROCESS_LAUNCHER)
	case TID_CACHE:
		return 1 == C.cef_currently_on(C.TID_CACHE)
	case TID_IO:
		return 1 == C.cef_currently_on(C.TID_IO)
	case TID_RENDERER:
		return 1 == C.cef_currently_on(C.TID_RENDERER)
	default:
		log.Warn("Unknown Thread ID type : ", id)
		return false
	}
}

const (
	TID_UI ThreadId = iota
	TID_DB
	TID_FILE
	TID_FILE_USER_BLOCKING
	TID_PROCESS_LAUNCHER
	TID_CACHE
	TID_IO
	TID_RENDERER
)

var (
	taskMap      = make(map[unsafe.Pointer]TaskToExecute)
	postTaskLock sync.Mutex
)

type TaskToExecute func()

//export go_TaskExecute
func go_TaskExecute(self *C.struct__cef_task_t) {
	postTaskLock.Lock()
	toExecute, ok := taskMap[unsafe.Pointer(self)]
	if ok {
		delete(taskMap, unsafe.Pointer(self))
	}
	postTaskLock.Unlock()
	if ok {
		toExecute()
	}
}

// allows you to execute a task on the specified thread
func PostTask(thread ThreadId, t TaskToExecute) {
	taskT := (*C.struct__cef_task_t)(
		C.calloc(1, C.sizeof_struct__cef_task_t))
	C.initialize_task(taskT)
	go_AddRef(unsafe.Pointer(taskT))

	// not defering the unlock, since i don't know if cef might immedialty execute it,
	// and thus we deadlock ourselves.
	postTaskLock.Lock()
	taskMap[unsafe.Pointer(taskT)] = t
	postTaskLock.Unlock()

	C.helper_cef_post_task(C.cef_thread_id_t(thread), taskT)
}
