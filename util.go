package chrome

import "C"
import (
	log "github.com/cihub/seelog"
)

//export goDebugLog
func goDebugLog(toLog *C.char) {
	log.Info("[C file] ", C.GoString(toLog))
}
