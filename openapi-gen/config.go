package main

import (
	"os"

	"github.com/iancoleman/strcase"
	"gopkg.in/yaml.v3"
)

// used to specify target specific patameters for code generation
type Config struct {
	AdditionalParameters map[string]any `yaml:"additionalParameters" json:"additionalParameters"`
	// FileNameModifier is function name to be used for generating the file name.
	FileNameModifier string `yaml:"fileNameModifier" json:"fileNameModifier"`
	// prefix imports with this string
	// directory where the generated client will be saved
	ClientOutputDirectory string `yaml:"clientOutputDirectory" json:"clientOutputDirectory"`
	FileExtension        string   `yaml:"fileExtension" json:"fileExtension"`
	CopyAdditionalFiles  []string `yaml:"copyAdditionalFiles" json:"copyAdditionalFiles"`
	GenerateRequestTypes bool     `yaml:"generateRequestTypes" json:"generateRequestTypes"`

	// relative path to the directory where the generated models will be saved
	ModelsSubpackage string `yaml:"modelsSubpackage" json:"modelsSubpackage"`
	// files to be copied to the models directory
	ModelsCopyFiles []string `yaml:"modelsCopyFiles" json:"modelsCopyFiles"`
}

func (c Config) getNameModifier() func(string) string {
	switch c.FileNameModifier {
	case "camelCase":
		return strcase.ToCamel
	case "snakeCase":
		return strcase.ToSnake
	}
	return func(s string) string { return s }
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
