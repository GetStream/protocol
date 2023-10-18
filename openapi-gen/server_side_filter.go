package main

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
)

const ServerSideKey = "x-stream-server-side"

var allHTTPMethods = []string{
	http.MethodConnect, http.MethodDelete, http.MethodGet,
	http.MethodHead, http.MethodOptions, http.MethodPatch,
	http.MethodPost, http.MethodPut, http.MethodTrace,
}

// removeServerSide removes the server side properties and operations from spec.
func removeServerSide(spec *openapi3.T) {
	spec.Paths = removeServerSideOperations(spec.Paths)
	if spec.Components != nil {
		spec.Components.Schemas = removeServerSideModels(spec.Components.Schemas)
	}
}

func nulifyOperatioin(item *openapi3.PathItem, method string) {
	switch method {
	case http.MethodConnect:
		item.Connect = nil
	case http.MethodDelete:
		item.Delete = nil
	case http.MethodGet:
		item.Get = nil
	case http.MethodHead:
		item.Head = nil
	case http.MethodOptions:
		item.Options = nil
	case http.MethodPatch:
		item.Patch = nil
	case http.MethodPost:
		item.Post = nil
	case http.MethodPut:
		item.Put = nil
	case http.MethodTrace:
		item.Trace = nil
	}
}

func removeServerSideOperations(paths openapi3.Paths) openapi3.Paths {
	for k, item := range paths {
		for _, method := range allHTTPMethods {
			op := item.GetOperation(method)
			if op == nil {
				continue
			}

			if op.Parameters != nil {
				newParams := make(openapi3.Parameters, 0, len(op.Parameters))
				for _, param := range op.Parameters {
					if param.Value == nil || param.Value.Extensions == nil || param.Value.Extensions[ServerSideKey] == nil {
						newParams = append(newParams, param)
					}
				}
				op.Parameters = newParams
			}

			if op.Extensions == nil || op.Extensions[ServerSideKey] == nil {
				continue
			}
			serverSide, ok := op.Extensions[ServerSideKey].(bool)
			if !ok || !serverSide {
				continue
			}
			nulifyOperatioin(item, method)
		}
		if len(item.Operations()) == 0 {
			delete(paths, k)
		}
	}
	return paths
}

func removeServerSideModels(schemas openapi3.Schemas) openapi3.Schemas {
	for k, item := range schemas {
		if item == nil || item.Value == nil {
			continue
		}
		removeServerSideProperties(item)

		if item.Value.Extensions == nil || item.Value.Extensions[ServerSideKey] == nil {
			continue
		}

		serverSide, ok := item.Value.Extensions[ServerSideKey].(bool)
		if !ok || !serverSide {
			continue
		}
		delete(schemas, k)
	}
	return schemas
}

func removeServerSideProperties(schema *openapi3.SchemaRef) {
	if schema.Value == nil || schema.Value.Properties == nil {
		return
	}
	for k, item := range schema.Value.Properties {
		if item.Value == nil {
			continue
		}

		removeServerSideProperties(item)
		if item.Value.Extensions == nil || item.Value.Extensions[ServerSideKey] == nil {
			continue
		}

		serverSide, ok := item.Value.Extensions[ServerSideKey].(bool)
		if !ok || !serverSide {
			continue
		}
		delete(schema.Value.Properties, k)
	}
}
