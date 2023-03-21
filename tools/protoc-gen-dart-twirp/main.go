package main

import (
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	gengo "google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo"
	"google.golang.org/protobuf/compiler/protogen"
)

//go:embed template.godart
var twirpTemplate string

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = gengo.SupportedFeatures
		for _, f := range gen.Files {
			if !f.Generate || len(f.Services) == 0 {
				continue
			}
			return generateFile(gen, f)
		}
		
		return nil
	})
}

type twirpFileData struct {
	GeneratedFileName string
	ProtoName         string
	Imports           []string
	Services          []*protogen.Service
}

func generateFile(gen *protogen.Plugin, file *protogen.File) error {
	// setup data struct and extract values given from protoc
	data := twirpFileData{}
	data.GeneratedFileName = file.GeneratedFilenamePrefix
	nameSplit := strings.Split(file.Proto.GetName(), "/")
	data.ProtoName = strings.TrimSuffix(nameSplit[len(nameSplit)-1], ".proto")
	// data.Imports = file.Proto.Dependency
	data.Imports = make([]string, file.Desc.Imports().Len())
	for i := 0; i < file.Desc.Imports().Len(); i++ {
		data.Imports[i] = strings.ReplaceAll(
			file.Desc.Imports().Get(i).Path(),
			".proto", "",
		)
	}
	data.Services = make([]*protogen.Service, len(file.Services))
	copy(data.Services, file.Services)

	// parse template
	protoTemplate, err := template.New(data.ProtoName).Funcs(funcMap).Parse(twirpTemplate)
	if err != nil {
		gen.Error(fmt.Errorf("error parsing template: %v", err))
		return err
	}

	// setup the GenerateFile and execute the template
	filename := file.GeneratedFilenamePrefix + ".pbtwirp.dart"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	err = protoTemplate.Execute(g, &data)
	if err != nil {
		gen.Error(fmt.Errorf("error executing template: %v", err))
		return err
	}

	return nil
}

func lowerFirstLetter(s string) string {
	return strings.ToLower(string(s[0])) + s[1:]
}

func upperFirstLetter(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

func removeNewLine(s string) string {
	return strings.ReplaceAll(s, "\n", "")
}

var funcMap = map[string]interface{}{
	"lowerFirstLetter": lowerFirstLetter,
	"upperFirstLetter": upperFirstLetter,
	"removeNewLine":    removeNewLine,
}
