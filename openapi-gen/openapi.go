package main

import (
	"context"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path"
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
	templates *template.Template
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

	tmpl := template.New("")
	tmpl.Funcs(template.FuncMap{
		"refToName": refToName,
		"toSnake":   strcase.ToSnake,
		"toCamel":   strcase.ToCamel,
		"has": func(sl []string, str string) bool {
			for _, s := range sl {
				if s == str {
					return true
				}
			}
			return false
		},
	})

	templates, err := tmpl.ParseFS(files, "*.tmpl")
	if err != nil {
		return nil, fmt.Errorf("parsing templates %w", err)
	}

	tl := &TemplateLoader{
		templates: templates,
	}

	return tl, nil
}

func (t *TemplateLoader) LoadTemplate(kind templateKind) *template.Template {
	return t.templates.Lookup(string(kind) + ".tmpl")
}

type TypeContext struct {
	Schema     *openapi3.Schema
	Name       string
	Additional map[string]any
	References []string
}

type RequestContext struct {
	Name       string
	Additional map[string]any
	References []string

	Parameters openapi3.Parameters
	Body       *openapi3.RequestBodyRef
}

// go run . -i ../openapi/video-openapi.yaml -o ./go-generated -l go
func main() {
	inputFile := flag.String("i", "", "yaml file to use for generating code")
	outputDir := flag.String("o", "", "output directory of generated code")
	targetLanguage := flag.String("l", "python", "target language to generate code for")
	configFile := flag.String("c", `config.yaml`, "config file to use for generating code")
	flag.Parse()

	if *inputFile == "" || *outputDir == "" {
		fmt.Println("must provide input file and output directory")
		os.Exit(1)
	}

	config, err := parseConfig(path.Join("templates", *targetLanguage, *configFile))
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

	templateLoader, err := NewTemplateLoader(*targetLanguage, true)
	if err != nil {
		fmt.Println("error loading template loader", err)
		os.Exit(1)
	}

	tmpl := templateLoader.LoadTemplate(TypeTemplate)

	for name, schema := range doc.Components.Schemas {
		ext := config.FileExtension
		f, err := os.Create(*outputDir + "/" + config.getNameModifier()(name) + ext)
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

	for pathName, path := range doc.Paths {
		for method, operation := range path.Operations() {
			if len(operation.Parameters) == 0 && operation.RequestBody == nil {
				continue
			}

			name := operation.OperationID

			if name == "" {
				// TODO: generate a name based on path and method
				fmt.Println("operationID is required for operation", pathName, method)
				os.Exit(1)
			}

			name = name + "Request"

			tmpl := templateLoader.LoadTemplate(RequestTemplate)

			ext := config.FileExtension
			f, err := os.Create(*outputDir + "/" + config.getNameModifier()(name) + "_" + method + ext)
			if err != nil {
				fmt.Println("error creating file", err)
				os.Exit(1)
			}
			defer f.Close()

			err = tmpl.Execute(f, RequestContext{
				Name:       name,
				Additional: config.AdditionalParameters,
				References: collectReferncesForRequest(operation.Parameters, operation.RequestBody),
				Parameters: operation.Parameters,
				Body:       operation.RequestBody,
			})

			if err != nil {
				fmt.Println("error executing template", err)
				os.Exit(1)
			}
		}
	}
}

func collectReferncesForRequest(parameters openapi3.Parameters, body *openapi3.RequestBodyRef) []string {
	var refs []string
	for _, param := range parameters {
		if param.Ref != "" {
			refs = append(refs, param.Ref)
			continue
		}

		if param.Value != nil {
			if param.Value.Schema != nil && param.Value.Schema.Ref != "" {
				refs = append(refs, param.Value.Schema.Ref)
			}
			if param.Value.Content != nil {
				for _, content := range param.Value.Content {
					if content.Schema != nil && content.Schema.Ref != "" {
						refs = append(refs, content.Schema.Ref)
					}
				}
			}
		}
	}

	if body != nil && body.Ref != "" {
		refs = append(refs, body.Ref)
	}

	if body != nil && body.Value != nil {
		if body.Value.Content != nil {
			for _, content := range body.Value.Content {
				if content.Schema != nil && content.Schema.Ref != "" {
					refs = append(refs, content.Schema.Ref)
				}
			}
		}
	}

	return refs
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
	if schema.OneOf != nil {
		for _, oneOf := range schema.OneOf {
			if oneOf.Ref != "" {
				refs = append(refs, oneOf.Ref)
			} else {
				refs = append(refs, getReferencesFromTypes(oneOf.Value)...)
			}
		}
	}

	for _, prop := range schema.Properties {
		if prop.Ref != "" {
			refs = append(refs, prop.Ref)
		} else if prop.Value != nil {
			refs = append(refs, getReferencesFromTypes(prop.Value)...)
		}
	}

	return refs
}
