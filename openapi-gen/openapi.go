package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
)

func GetTemplates(language string) (string, error) {
	return "", nil
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
