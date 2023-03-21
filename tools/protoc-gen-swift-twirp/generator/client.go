package generator

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
	"unicode"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

//go:embed template.goswift
var apiTemplate string

type Service struct {
	Name    string
	Package string
	Methods []ServiceMethod
}

func (s Service) ClassName() string {
	return toSnake(s.Package) + "_" + s.Name
}

type ServiceMethod struct {
	Name       string
	Path       string
	InputArg   string
	InputType  string
	OutputType string
}

func NewAPIContext() APIContext {
	ctx := APIContext{}

	return ctx
}

type APIContext struct {
	Services []*Service
	Imports  []Import
}

type Import struct {
	Path string
}

func (ctx *APIContext) ApplyImports(d *descriptorpb.FileDescriptorProto) {
	var deps []Import

	if len(ctx.Services) > 0 {
		deps = append(deps, Import{"Foundation"})
		deps = append(deps, Import{"SwiftProtobuf"})
	}

	for _, dep := range d.Dependency {
		importPath := path.Dir(dep)
		sourceDir := path.Dir(*d.Name)
		sourceComponents := strings.Split(sourceDir, fmt.Sprintf("%c", os.PathSeparator))
		distanceFromRoot := len(sourceComponents)
		for _, pathComponent := range sourceComponents {
			if strings.HasPrefix(importPath, pathComponent) {
				importPath = strings.TrimPrefix(importPath, pathComponent)
				distanceFromRoot--
			}
		}
		fileName := swiftFileName(dep)
		fullPath := fileName
		fullPath = path.Join(importPath, fullPath)
		if distanceFromRoot > 0 {
			for i := 0; i < distanceFromRoot; i++ {
				fullPath = path.Join("..", fullPath)
			}
		}
	}
	ctx.Imports = deps
}

func CreateClientAPI(gen *protogen.Plugin, f *protogen.File) error {
	ctx := NewAPIContext()
	pkg := f.Proto.GetPackage()

	// Parse all Services for generating typescript method interfaces and default client implementations
	for _, s := range f.Proto.GetService() {
		service := &Service{
			Name:    s.GetName(),
			Package: pkg,
		}

		for _, m := range s.GetMethod() {
			methodPath := m.GetName()
			methodName := strings.ToLower(methodPath[0:1]) + methodPath[1:]
			in := removePkg(m.GetInputType())
			arg := strings.ToLower(in[0:1]) + in[1:]
			in = toSnake(m.GetInputType()[1:])
			method := ServiceMethod{
				Name:       methodName,
				Path:       methodPath,
				InputArg:   arg,
				InputType:  in,
				OutputType: toSnake(m.GetOutputType()[1:]),
			}

			service.Methods = append(service.Methods, method)
		}
		ctx.Services = append(ctx.Services, service)
	}
	ctx.ApplyImports(f.Proto)
	funcMap := template.FuncMap{}

	t, err := template.New("client_api").Funcs(funcMap).Parse(apiTemplate)
	if err != nil {
		return err
	}

	b := bytes.NewBufferString("")
	err = t.Execute(b, ctx)
	if err != nil {
		return err
	}

	file := gen.NewGeneratedFile(f.GeneratedFilenamePrefix+".twirp.swift", f.GoImportPath)
	_, err = file.Write(b.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func toSnake(s string) string {
	ss := strings.Split(s, ".")

	res := []string{}
	for _, v := range ss {
		parts := strings.Split(v, "_")
		var subParts []string
		for _, p := range parts {
			subParts = append(subParts, string(byte(unicode.ToUpper(rune(p[0]))))+p[1:])
		}
		// first to upper
		res = append(res, strings.Join(subParts, ""))
	}
	return strings.Join(res, "_")
}

func removePkg(s string) string {
	p := strings.Split(s, ".")
	return p[len(p)-1]
}
