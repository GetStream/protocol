package main

import (
	"os"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestGenerate(t *testing.T) {
	loader, err := NewTemplateLoader("go", false)
	if err != nil {
		t.Fatal(err)
	}

	tmpl, err := loader.LoadTemplate(TypeTemplate)
	if err != nil {
		t.Fatal(err)
	}

	context := TypeContext{
		Schema: &openapi3.Schema{
			Type: "object",
			Properties: map[string]*openapi3.SchemaRef{
				"text": {
					Value: &openapi3.Schema{
						Type: "string",
					},
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
