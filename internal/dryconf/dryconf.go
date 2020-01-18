package dryconf

import (
	"os"

	"github.com/dmatusiewicz/dryconf/internal/flags"
	"github.com/dmatusiewicz/dryconf/internal/logs"
	"github.com/dmatusiewicz/dryconf/internal/version"
	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func Generate() {
	flags.Parse()
	logger := logs.Configure(flags.Debug())
	if flags.Version() {
		logger.Print("Printing version.")
		version.Print()
		os.Exit(0)
	}

}
