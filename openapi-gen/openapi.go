package main

import (
	"context"
	"flag"
	"fmt"
	"io"
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
		"successfulResponse": func(responses openapi3.Responses) *openapi3.SchemaRef {
			for code, response := range responses {
				if code == "200" || code == "201" {
					return response.Value.Content["application/json"].Schema
				}
			}
			return nil
		},
		"requestSchema": func(operation openapi3.Operation) *openapi3.SchemaRef {
			if operation.RequestBody == nil {
				return nil
			}
			return operation.RequestBody.Value.Content["application/json"].Schema
		},
		"operationContext":   operationContext,
		"sortedProperties":   sortedProperties,
		"requiredParameters": requiredParameters,
		"optionalParameters": optionalParameters,
		"clientReferences":   clientReferences,
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

	for _, fileName := range config.CopyAdditionalFiles {
		dst, err := os.Create(*outputDir + "/" + fileName)
		if err != nil {
			fmt.Println("error opening additional file", err)
			os.Exit(1)
		}

		src, err := os.Open("templates/" + *targetLanguage + "/" + fileName)
		if err != nil {
			fmt.Println("error opening additional file", err)
			os.Exit(1)
		}

		_, err = io.Copy(dst, src)
		if err != nil {
			fmt.Println("error copying additional file", err)
			os.Exit(1)
		}

		dst.Close()
		src.Close()
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
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if config.GenerateRequestTypes {
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
				f, err := os.Create(*outputDir + "/" + config.getNameModifier()(name) + "_" + strings.ToLower(method) + ext)
				if err != nil {
					fmt.Println("error creating file", err)
					os.Exit(1)
				}
				//we shouldn't defer here because we want to close the file after each request generation
				defer f.Close()

				err = tmpl.Execute(f, RequestContext{
					Name:       name,
					Additional: config.AdditionalParameters,
					References: collectReferncesForRequest(operation),
					Parameters: operation.Parameters,
					Body:       operation.RequestBody,
				})

				if err != nil {
					fmt.Println("error generating", name, err)
					os.Exit(1)
				}
			}
		}
	}

	f, err := os.Create(path.Join(*outputDir, "client"+config.FileExtension))
	if err != nil {
		fmt.Println("error creating file", err)
		os.Exit(1)
	}
	defer f.Close()

	tp := templateLoader.LoadTemplate("client")
	if tp != nil {
		err = tp.Execute(f, doc)
		if err != nil {
			fmt.Println("error generating client", err)
			os.Exit(1)
		}
	}
}

func collectReferncesForRequest(operation *openapi3.Operation) []string {
	set := make(map[string]struct{})
	buildRequestReferencesSet(operation, set)

	refs := make([]string, 0, len(set))
	for ref := range set {
		refs = append(refs, ref)
	}

	return refs
}

func buildRequestReferencesSet(operation *openapi3.Operation, refs map[string]struct{}) {
	for _, param := range operation.Parameters {
		if param.Ref != "" {
			refs[param.Ref] = struct{}{}
			continue
		}

		if param.Value != nil {
			if param.Value.Schema != nil && param.Value.Schema.Ref != "" {
				refs[param.Value.Schema.Ref] = struct{}{}
			}
			if param.Value.Content != nil {
				for _, content := range param.Value.Content {
					if content.Schema != nil && content.Schema.Ref != "" {
						refs[content.Schema.Ref] = struct{}{}
					}
				}
			}
		}
	}

	body := operation.RequestBody
	if body != nil && body.Ref != "" {
		refs[body.Ref] = struct{}{}
	}

	if body != nil && body.Value != nil {
		if body.Value.Content != nil {
			for _, content := range body.Value.Content {
				if content.Schema != nil && content.Schema.Ref != "" {
					refs[content.Schema.Ref] = struct{}{}
				}
			}
		}
	}

	if operation.Responses != nil {
		for _, response := range operation.Responses {
			if response.Value.Content != nil {
				for _, content := range response.Value.Content {
					if content.Schema != nil && content.Schema.Ref != "" {
						refs[content.Schema.Ref] = struct{}{}
					}
				}
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
