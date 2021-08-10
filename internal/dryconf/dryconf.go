package dryconf

import (
	"os"

	"github.com/dmatusiewicz/dryconf/internal/configuration"
	"github.com/dmatusiewicz/dryconf/internal/flags"
	"github.com/dmatusiewicz/dryconf/internal/logs"
	"github.com/dmatusiewicz/dryconf/internal/version"
)

// Log costam

// Generate core function that will generate the configuraion file. This is the place where it all starts
func Generate() {

	flags.Parse()

	loggerS := logs.Configure(flags.Debug())

	if flags.Version() {
		version.Print()
		os.Exit(0)
	}

	// yamlmerger.MergeWrapper(loggerS)

	err := configuration.Load(loggerS, flags.Config())
	if err != nil {
		loggerS.Fatal().Err(err).Msg("Error when loading configuration.")
	}

}
