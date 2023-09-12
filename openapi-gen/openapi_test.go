package main

import (
	"os"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestGenerate(t *testing.T) {
	loader, err := NewTemplateLoader("python", false, PrepareBuiltinFunctions(&Config{}))
	if err != nil {
		t.Fatal(err)
	}

	tmpl := loader.LoadTemplate(TypeTemplate)

	context := TypeContext{
		Schema: &openapi3.Schema{
			Type: "object",
			Properties: map[string]*openapi3.SchemaRef{
				"text": {
					Value: &openapi3.Schema{
						Type: "string",
					},
				},
				"arr": {
					Value: &openapi3.Schema{
						Type: "array",
						Items: &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type: "string",
							},
						},
					},
				},
				"othertype": {
					Ref: "#/components/schemas/OtherType",
				},
			},
		},
		Name: "Message",
	}

	out := os.Stdout
	err = tmpl.Execute(out, context)
	if err != nil {
		t.Fatal(err)
	}

}
