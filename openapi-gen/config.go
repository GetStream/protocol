package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

type ParameterType string

const (
	ParameterTypeBool   ParameterType = "bool"
	ParameterTypeInt    ParameterType = "int"
	ParameterTypeString ParameterType = "string"
)

type AdditionalParameterSpec struct {
	Usage    string        `yaml:"usage" json:"usage"`
	Type     ParameterType `yaml:"type" json:"type"`
	Default  interface{}   `yaml:"default" json:"default"`
	Required bool          `yaml:"required" json:"required"`
}

// used to specify target specific patameters for code generation
type TargetConfig struct {
	AdditionalParameters map[string]AdditionalParameterSpec `yaml:"additionalParameters" json:"additionalParameters"`
}

// put target specific paramters into flagSet and return context map
func getParamsFromEnv(parameters map[string]AdditionalParameterSpec) (map[string]any, error) {
	params := make(map[string]any)
	for key, v := range parameters {
		var err error
		rawparam := os.Getenv(key)

		if rawparam == "" {
			if v.Required {
				return nil, fmt.Errorf("required parameter %s not set", key)
			}
			params[key] = v.Default
			continue
		}

		switch v.Type {
		case ParameterTypeBool:
			params[key], err = strconv.ParseBool(rawparam)
		case ParameterTypeInt:
			params[key], err = strconv.Atoi(rawparam)
		case ParameterTypeString:
			params[key] = rawparam
		default:
			return nil, fmt.Errorf("unknown parameter type %s", v.Type)
		}

		if err != nil {
			return nil, fmt.Errorf("error parsing parameter %s as %s: %w", key, v.Type, err)
		}
	}
	return params, nil
}

func parseTargetConfig(reader io.Reader) (*TargetConfig, error) {
	yamlDecoder := yaml.NewDecoder(reader)
	var config TargetConfig
	err := yamlDecoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
