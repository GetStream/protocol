package main

import (
	"os"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestGenerateGo(t *testing.T) {
	loader, err := NewTemplateLoader("go", false)
	if err != nil {
		t.Fatal(err)
	}

	config := Config{
		AdditionalParameters: map[string]interface{}{
			"package": "stream",
		},
	}

	tmpl := loader.LoadTemplate(TypeTemplate)
	doc, err := openapi3.NewLoader().LoadFromData(
		[]byte(`openapi: 3.0.0
info:
  title: Stream API
  description: Stream API
  version: 1.0.0
components:
  schemas:
    MemberResponse:
      properties:
        created_at:
          description: Date/time of creation
          format: date-time
          title: Created at
          type: string
          x-stream-index: "005"
        custom:
          additionalProperties: {}
          description: Custom member response data
          title: Custom
          type: object
          x-stream-index: "004"
        deleted_at:
          description: Date/time of deletion
          format: date-time
          title: Deleted at
          type: string
          x-stream-index: "007"
        role:
          title: Role
          type: string
          x-stream-index: "003"
        updated_at:
          description: Date/time of the last update
          format: date-time
          title: Updated at
          type: string
          x-stream-index: "006"
        user:
          $ref: '#/components/schemas/UserResponse'
          x-stream-index: "001"
        user_id:
          type: string
          x-stream-index: "002"
      required:
      - user
      - user_id
      - custom
      - created_at
      - updated_at
      type: object
    UpdateCallMembersResponse:
      nullable: true
      properties:
        duration:
          description: Duration of the request in human-readable format
          title: Duration
          type: string
          x-stream-index: "001.001"
        members:
          items:
            $ref: '#/components/schemas/MemberResponse'
          type: array
          x-stream-index: "002"
      required:
      - duration
      - members
      type: object   
    UserResponse:
      properties:
        created_at:
          description: Date/time of creation
          format: date-time
          title: Created at
          type: string
          x-stream-index: "002"
        custom:
          additionalProperties: {}
          type: object
          x-stream-index: "001.004"
        deleted_at:
          description: Date/time of deletion
          format: date-time
          title: Deleted at
          type: string
          x-stream-index: "004"
        id:
          type: string
          x-stream-index: "001.001"
        image:
          type: string
          x-stream-index: "001.003"
        name:
          type: string
          x-stream-index: "001.002"
        role:
          type: string
          x-stream-index: "001.005"
        teams:
          items:
            type: string
          type: array
          x-stream-index: "001.006"
        updated_at:
          description: Date/time of the last update
          format: date-time
          title: Updated at
          type: string
          x-stream-index: "003"
      required:
      - id
      - custom
      - role
      - teams
      - created_at
      - updated_at
      type: object
  `))
	if err != nil {
		t.Fatal(err)
	}
	out := os.Stdout
	for name, schema := range doc.Components.Schemas {

		err = tmpl.Execute(out, TypeContext{
			Name:       name,
			Schema:     schema.Value,
			Additional: config.AdditionalParameters,
			References: getReferencesFromTypes(schema.Value),
		})
		if err != nil {
			t.Fatal(err)
		}
	}

}

