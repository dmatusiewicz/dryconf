package logs

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
)

var zl zerolog.Logger

func Configure(debug bool) zerolog.Logger {
	var once sync.Once
	once.Do(func() {

		if debug {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
		zl = zerolog.New(os.Stderr).With().Timestamp().Logger()

	})
	zl.Print("Logger has been configured.")
	return zl
}
