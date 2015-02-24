package chrome

import (
	log "github.com/cihub/seelog"
	"os"
)

func DisableLog() {
	logger = log.Disabled
	log.ReplaceLogger(logger)
}

func EnableLog() {
	logger, _ = log.LoggerFromWriterWithMinLevelAndFormat(os.Stdout, 0, "[%Level] %File:%Line: %Msg %n")
	log.ReplaceLogger(logger)
}
