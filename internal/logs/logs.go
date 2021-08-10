package logs

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
)

type Log interface {
	Printf(format string, v ...interface{})
	Print(v ...interface{})
	Fatal() *zerolog.Event
}

func Configure(debug bool) Log {
	var once sync.Once
	var zeroLogInterface Log
	once.Do(func() {

		if debug {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
		zeroLogObject := zerolog.New(os.Stderr).With().Timestamp().Logger()
		zeroLogInterface = &zeroLogObject

	})
	zeroLogInterface.Print("Logger has been configured.")
	return zeroLogInterface
}
