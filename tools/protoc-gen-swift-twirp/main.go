package main

import (
	gengo "google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo"
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/GetStream/video/tools/protoc-gen-swift-twirp/generator"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate || len(f.Services) == 0 {
				continue
			}
			err := generator.CreateClientAPI(gen, f)
			if err != nil {
				return err
			}
		}
		gen.SupportedFeatures = gengo.SupportedFeatures
		return nil
	})
}
