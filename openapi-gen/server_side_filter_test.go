package main

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func Test_removeServerSide(t *testing.T) {
	type args struct {
		spec *openapi3.T
	}
	tests := []struct {
		name   string
		args   args
		expect func(*testing.T, *openapi3.T)
	}{
		{
			name: "cleanOperations",
			args: args{
				spec: &openapi3.T{
					Paths: openapi3.Paths{
						"/test-server-side": &openapi3.PathItem{
							Get: &openapi3.Operation{
								Extensions: map[string]interface{}{
									ServerSideKey: true,
								},
							},
						},
						"/test": &openapi3.PathItem{
							Get: &openapi3.Operation{
								Parameters: openapi3.Parameters{
									&openapi3.ParameterRef{
										Value: &openapi3.Parameter{
											Name: "test",
											In:   "query",
											Extensions: map[string]interface{}{
												ServerSideKey: true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expect: func(t *testing.T, spec *openapi3.T) {
				if _, ok := spec.Paths["/test-server-side"]; ok {
					t.Error("removeServerSide does not delete server side path")
				}
				if _, ok := spec.Paths["/test"]; !ok || spec.Paths["/test"] == nil || spec.Paths["/test"].Get == nil {
					t.Error("removeServerSide delete path without server side")
				}
				if len(spec.Paths["/test"].Get.Parameters) != 0 {
					t.Error("removeServerSide does not delete server side parameters")
				}
			},
		},
		{
			name: "cleanSchemas",
			args: args{
				spec: &openapi3.T{
					Components: &openapi3.Components{
						Schemas: map[string]*openapi3.SchemaRef{
							"test": {
								Value: &openapi3.Schema{
									Extensions: map[string]interface{}{
										ServerSideKey: true,
									},
								},
							},
							"TestProperty": {
								Value: &openapi3.Schema{
									Properties: map[string]*openapi3.SchemaRef{
										"server": {
											Value: &openapi3.Schema{
												Type: "string",
												Extensions: map[string]interface{}{
													ServerSideKey: true,
												},
											},
										},
										"client": {
											Value: &openapi3.Schema{
												Type: "string",
											},
										},
										"Deep": {
											Value: &openapi3.Schema{
												Type: "object",
												Properties: map[string]*openapi3.SchemaRef{
													"server": {
														Value: &openapi3.Schema{
															Type: "string",
															Extensions: map[string]interface{}{
																ServerSideKey: true,
															},
														},
													},
													"client": {
														Value: &openapi3.Schema{
															Type: "string",
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expect: func(t *testing.T, spec *openapi3.T) {
				if _, ok := spec.Components.Schemas["test"]; ok {
					t.Error("removeServerSide does not delete server side schema")
				}
				if _, ok := spec.Components.Schemas["TestProperty"]; !ok || spec.Components.Schemas["TestProperty"] == nil || spec.Components.Schemas["TestProperty"].Value == nil {
					t.Error("removeServerSide delete schema without server side")
				}
				v := spec.Components.Schemas["TestProperty"].Value
				if _, ok := v.Properties["server"]; ok {
					t.Error("removeServerSide does not delete server side property")
				}
				if _, ok := v.Properties["client"]; !ok || v.Properties["client"] == nil || v.Properties["client"].Value == nil {
					t.Error("removeServerSide delete property without server side")
				}
				if _, ok := v.Properties["Deep"]; !ok || v.Properties["Deep"] == nil || v.Properties["Deep"].Value == nil {
					t.Error("removeServerSide delete property without server side")
				}
				dv := v.Properties["Deep"].Value
				if _, ok := dv.Properties["server"]; ok {
					t.Error("removeServerSide does not delete server side deep property")
				}
				if _, ok := dv.Properties["client"]; !ok || dv.Properties["client"] == nil || dv.Properties["client"].Value == nil {
					t.Error("removeServerSide delete deep property without server side")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			removeServerSide(tt.args.spec)
			if tt.expect != nil {
				tt.expect(t, tt.args.spec)
			}
		})
	}
}
