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
			name: "removeServerSide",
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
