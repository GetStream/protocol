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
			if op == nil || op.Extensions == nil || op.Extensions[ServerSideKey] == nil {
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
