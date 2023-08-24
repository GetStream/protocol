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
	"github.com/iancoleman/strcase"
)

type templateKind string

const (
	TypeTemplate    templateKind = "type"
	RequestTemplate templateKind = "request"
)

type TemplateLoader struct {
	files fs.FS
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
	}

	return tl, nil
}

func (t *TemplateLoader) LoadTemplate(kind templateKind) (*template.Template, error) {
	tmpl := template.New(string(kind) + ".tmpl")
	tmpl.Funcs(template.FuncMap{
		"refToName": refToName,
		"toSnake":   strcase.ToSnake,
		"toCamel":   strcase.ToCamel,
	})

	tmpl, err := tmpl.ParseFS(t.files, string(kind)+".tmpl")
	if err != nil {
		return nil, fmt.Errorf("parsing template %w", err)
	}

	return tmpl, nil
}

type TypeContext struct {
	Schema     *openapi3.Schema
	Name       string
	Additional map[string]any
	References []string
}

// go run . -i ../openapi/video-openapi.yaml
func main() {
	inputFile := flag.String("i", "", "yaml file to use for generating code")
	outputDir := flag.String("o", "", "output directory of generated code")
	configFile := flag.String("c", "config.yaml", "config file to use for generating code")
	flag.Parse()

	if *inputFile == "" || *outputDir == "" {
		fmt.Println("must provide input file and output directory")
		os.Exit(1)
	}

	config, err := parseConfig(*configFile)
	if err != nil {
		fmt.Println("error parsing config", err)
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

	templateLoader, err := NewTemplateLoader("python", true)
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

			f, err := os.Create(*outputDir + "/" + name + ".py")
			if err != nil {
				fmt.Println("error creating file", err)
				os.Exit(1)
			}
			defer f.Close()

			err = tmpl.Execute(f, TypeContext{
				Name:       name,
				Schema:     schema.Value,
				Additional: config.AdditionalParameters,
				References: getReferencesFromTypes(schema.Value),
			})

			if err != nil {
				fmt.Println("error executing template", err)
				os.Exit(1)
			}
		}
	}
}

func refToName(ref string) string {
	return strings.TrimPrefix(ref, "#/components/schemas/")
}

func getReferencesFromTypes(schema *openapi3.Schema) []string {
	if schema == nil {
		return nil
	}

	if schema.Type == "array" {
		if schema.Items.Ref != "" {
			return []string{schema.Items.Ref}
		}
		if schema.Items.Value != nil {
			return getReferencesFromTypes(schema.Items.Value)
		}
	}

	var refs []string
	for _, prop := range schema.Properties {
		if prop.Ref != "" {
			refs = append(refs, prop.Ref)
		} else if prop.Value != nil {
			refs = append(refs, getReferencesFromTypes(prop.Value)...)
		}
	}

	return refs
}
