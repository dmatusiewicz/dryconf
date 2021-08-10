package configuration

import (
	"github.com/dmatusiewicz/dryconf/internal/logs"
	"gopkg.in/yaml.v2"
)

// Configuration costam
type Configuration struct {
	ch map[string]interface{}
	co interface{}
}

var yc Configuration

var logger logs.Log

// Load costam
func Load(l logs.Log, config string) error {
	logger = l
	logger.Printf("Loading config file: %s", config)
	d, err := loadfile(config)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to load main configuration file")
	}
	err = yaml.Unmarshal(d, &yc.ch)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to Unmarshal the data from file")
	}

	err = yc.Parse()
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not find [config] section in the loaded configuration. It is obligatory to build configuration hierarchy.")
	}
	logger.Printf("File: %s has been read succesfully.", config)

	return nil
}
