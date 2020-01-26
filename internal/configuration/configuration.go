package configuration

import (
	"github.com/goccy/go-yaml"
	"github.com/rs/zerolog"
)

type Configuration struct {
	ch map[string]interface{}
	co interface{}
}

var yc Configuration

var logger zerolog.Logger

func Load(l zerolog.Logger, config string) error {
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
