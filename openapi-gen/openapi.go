package main

import (
	"context"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
)

type templateKind string

const (
	TypeTemplate    templateKind = "type"
	RequestTemplate templateKind = "request"
)

type TemplateLoader struct {
	files fs.FS

	envParams map[string]any
}

func NewTemplateLoader(targetLanguage string, useDisk bool) (*TemplateLoader, error) {
	var root fs.FS = templates
	if useDisk {
		root = os.DirFS(".")
	}

	files, err := fs.Sub(root, "templates/"+targetLanguage)
	if err != nil {
		return nil, fmt.Errorf("opening subdirectory fs %w, useDisk %v", err, useDisk)
	}

	tl := &TemplateLoader{
		files: files,
	}

	file, err := files.Open(`config.yaml`)
	if err == nil {
		defer file.Close()

		cfg, err := parseTargetConfig(file)
		if err != nil {
			return nil, fmt.Errorf("parsing config %w", err)
		}

		tl.envParams, err = getParamsFromEnv(cfg.AdditionalParameters)
		if err != nil {
			return nil, fmt.Errorf("getting params from env %w", err)
		}
	}

	return tl, nil
}

func (t *TemplateLoader) LoadTemplate(kind templateKind) (*template.Template, error) {
	tmpl, err := template.ParseFS(t.files, string(kind)+".tmpl")
	if err != nil {
		return nil, fmt.Errorf("parsing template %w", err)
	}

	return tmpl, nil
}

type TypeContext struct {
	Schema     *openapi3.Schema
	Name       string
	Additional map[string]any
}

// go run . -i ../openapi/video-openapi.yaml
func main() {
	inputFile := flag.String("i", "", "yaml file to use for generating code")
	outputDir := flag.String("o", "", "output directory of generated code")
	flag.Parse()

	if *inputFile == "" || *outputDir == "" {
		fmt.Println("must provide input file and output directory")
		os.Exit(1)
	}

	os.Mkdir(*outputDir, 0755)

	doc, err := openapi3.NewLoader().LoadFromFile(*inputFile)
	if err != nil {
		fmt.Println("error loading file", err)
		os.Exit(1)
	}

	err = doc.Validate(context.Background())
	if err != nil {
		fmt.Println("error validating doc", err)
		os.Exit(1)
	}

	templateLoader, err := NewTemplateLoader("go", true)
	if err != nil {
		fmt.Println("error loading template loader", err)
		os.Exit(1)
	}

	for name, schema := range doc.Components.Schemas {
		if strings.HasSuffix(name, "Request") {
			tmpl, err := templateLoader.LoadTemplate(TypeTemplate)
			if err != nil {
				fmt.Println("error loading template", err)
				os.Exit(1)
			}

			f, err := os.Create(*outputDir + "/" + name + ".go")
			if err != nil {
				fmt.Println("error creating file", err)
				os.Exit(1)
			}
			defer f.Close()

			err = tmpl.Execute(f, TypeContext{
				Name:       name,
				Schema:     schema.Value,
				Additional: templateLoader.envParams,
			})

			if err != nil {
				fmt.Println("error executing template", err)
				os.Exit(1)
			}
		}
	}
}
