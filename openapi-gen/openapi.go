package openapigen

import "github.com/getkin/kin-openapi/openapi3"

func ParseOpenAPISpec(filepath string) (*openapi3.T, error) {
	return openapi3.NewLoader().LoadFromFile(filepath)
}

func GetTemplates(language string) (string, error) {
	return "", nil
}
