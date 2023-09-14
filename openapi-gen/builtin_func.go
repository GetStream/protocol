package main

import (
	"strings"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/iancoleman/strcase"
)

func PrepareBuiltinFunctions(config *Config) template.FuncMap {
	return template.FuncMap{
		"refToName": refToName,
		"toSnake":   strcase.ToSnake,
		"toCamel":   strcase.ToCamel,
		"lower":     strings.ToLower,
		"has": func(sl []string, str string) bool {
			for _, s := range sl {
				if s == str {
					return true
				}
			}
			return false
		},
		"successfulResponse": func(responses openapi3.Responses) *openapi3.SchemaRef {
			for code, response := range responses {
				if code == "200" || code == "201" {
					return response.Value.Content["application/json"].Schema
				}
			}
			return nil
		},
		"requestSchema": func(operation *openapi3.Operation) *openapi3.SchemaRef {
			if operation.RequestBody == nil {
				return nil
			}

			if operation.RequestBody.Value.Content["application/json"] == nil {
				return nil
			}

			return operation.RequestBody.Value.Content["application/json"].Schema
		},
		"operationContext":   operationContext,
		"sortedProperties":   sortedProperties,
		"requiredParameters": requiredParameters,
		"optionalParameters": optionalParameters,
		"clientReferences":   clientReferences,
		"additionalParameters": func() map[string]any {
			return config.AdditionalParameters
		},
	}
}

type OperationWithPathAndMethod struct {
	Path   string
	Method string
	*openapi3.Operation
}

func operationContext(operation *openapi3.Operation, method, path string) *OperationWithPathAndMethod {
	return &OperationWithPathAndMethod{
		Path:      path,
		Method:    method,
		Operation: operation,
	}
}

func requiredParameters(parameters openapi3.Parameters) openapi3.Parameters {
	var required openapi3.Parameters
	for _, param := range parameters {
		if param.Value.Required {
			required = append(required, param)
		}
	}
	return required
}

func optionalParameters(parameters openapi3.Parameters) openapi3.Parameters {
	var optional openapi3.Parameters
	for _, param := range parameters {
		if !param.Value.Required {
			optional = append(optional, param)
		}
	}
	return optional
}

func clientReferences(paths openapi3.Paths) []string {
	set := make(map[string]struct{})
	for _, path := range paths {
		for _, operation := range path.Operations() {
			buildRequestReferencesSet(operation, set)
		}
	}

	refs := make([]string, 0, len(set))
	for ref := range set {
		refs = append(refs, ref)
	}

	return refs
}

type SchemaRefWithName struct {
	Name     string
	Required bool
	*openapi3.SchemaRef
}

func sortedProperties(schema openapi3.Schema) []SchemaRefWithName {
	var requiredNames = make(map[string]struct{})
	var required, optional []SchemaRefWithName

	for _, name := range schema.Required {
		requiredNames[name] = struct{}{}
	}

	for name, prop := range schema.Properties {
		if _, ok := requiredNames[name]; ok {
			required = append(required, SchemaRefWithName{
				Name:      name,
				Required:  true,
				SchemaRef: prop,
			})
		} else {
			optional = append(optional, SchemaRefWithName{
				Name:      name,
				Required:  false,
				SchemaRef: prop,
			})
		}
	}
	return append(required, optional...)
}

func refToName(ref string) string {
	return strings.TrimPrefix(ref, "#/components/schemas/")
}
