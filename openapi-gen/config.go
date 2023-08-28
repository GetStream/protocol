package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

// used to specify target specific patameters for code generation
type Config struct {
	AdditionalParameters map[string]any `yaml:"additionalParameters" json:"additionalParameters"`
}

func parseConfig(filePath string) (*Config, error) {
	reader, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer reader.Close()

	yamlDecoder := yaml.NewDecoder(reader)
	var config Config
	err = yamlDecoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
