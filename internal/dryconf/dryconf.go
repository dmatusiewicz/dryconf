package dryconf

import (
	"github.com/dmatusiewicz/dryconf/internal/flags"
	"github.com/dmatusiewicz/dryconf/internal/logs"
)

func Generate() {
	flags.Parse()

	logger := logs.Configure(flags.Debug())
	logger.Print("Logger has been configured.")
	logger.Print("Logger has been configured.")

	if flags.Version() {

	}

}
