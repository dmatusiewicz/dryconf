package dryconf

import (
	"os"

	"github.com/dmatusiewicz/dryconf/internal/configuration"
	"github.com/dmatusiewicz/dryconf/internal/flags"
	"github.com/dmatusiewicz/dryconf/internal/logs"
	"github.com/dmatusiewicz/dryconf/internal/version"
	"github.com/rs/zerolog"
)

var logger zerolog.Logger

// Generate core function that will generate the configuraion file. This is the place where it all starts
func Generate() {

	flags.Parse()

	logger := logs.Configure(flags.Debug())

	if flags.Version() {
		version.Print()
		os.Exit(0)
	}

	err := configuration.Load(logger, flags.Config())
	if err != nil {
		logger.Fatal().Err(err).Msg("Error when loading configuration.")
	}

}
