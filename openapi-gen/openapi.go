package main

import (
	"context"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
)

type templateKind string

const (
	TypeTemplate templateKind = "type"
)

type TemplateLoader struct {
	files fs.FS
}

func NewTemplateLoader(target string, useDisk bool) (*TemplateLoader, error) {
	var root fs.FS = templates
	if useDisk {
		root = os.DirFS(".")
	}

	files, err := fs.Sub(root, "templates/"+target)
	if err != nil {
		return nil, fmt.Errorf("opening subdirectory fs %w, useDisk %v", err, useDisk)
	}

	return &TemplateLoader{
		files: files,
	}, nil
}

func (t *TemplateLoader) LoadTemplate(kind templateKind) (*template.Template, error) {
	return template.ParseFS(t.files, string(kind)+".tmpl")
}

type TypeContext struct {
	Schema *openapi3.Schema
	Name   string
}

// go run . -i ../openapi/video-openapi.yaml
func main() {
	inputFile := flag.String("i", "", "yaml file to use for generating code")
	flag.Parse()

	if *inputFile == "" {
		fmt.Println("must provide input file")
		os.Exit(1)
	}

	doc, err := openapi3.NewLoader().LoadFromFile(*inputFile)
	if err != nil {
		panic(err)
	}

	err = doc.Validate(context.Background())
	if err != nil {
		panic(err)
	}

	for name, schema := range doc.Components.Schemas {
		fmt.Println(name, schema.Value.Type, schema.Value.Format)
	}
}
