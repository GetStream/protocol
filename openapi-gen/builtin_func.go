package main

import (
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/iancoleman/strcase"
	"golang.org/x/exp/slices"
)

type set map[string]struct{}

func newSet() set {
	return make(set)
}

func (s set) Add(value string) bool {
	if _, ok := s[value]; ok {
		return false
	}
	s[value] = struct{}{}
	return true
}

func (s set) Contains(value string) bool {
	_, ok := s[value]
	return ok
}

func toConstant(s string) string {
	return strings.ToUpper(strings.ReplaceAll(s, "-", "_"))
}

// filter returns a new slice holding only
// the elements of s that satisfy f()
func Filter(vs []SchemaRefWithName, f func(SchemaRefWithName) bool) []SchemaRefWithName {
	vsf := make([]SchemaRefWithName, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func FilterRequired(vs []SchemaRefWithName) []SchemaRefWithName {
	return Filter(vs, func(v SchemaRefWithName) bool { return v.Required })
}

func FilterNotRequired(vs []SchemaRefWithName) []SchemaRefWithName {
	return Filter(vs, func(v SchemaRefWithName) bool { return !v.Required })
}

func PrepareBuiltinFunctions(config *Config) template.FuncMap {
	return template.FuncMap{
		"refToName": refToName,
		"toSnake":   strcase.ToSnake,
		"toCamel":   strcase.ToCamel,
		"lower":     strings.ToLower,
		"newSet":    newSet,
		"append": func(value []string, elems ...string) []string {
			return append(value, elems...)
		},
		"list": func(value ...string) []string {
			return value
		},
		"contains": func(value []string, elem string) bool {
			for _, el := range value {
				if el == elem {
					return true
				}
			}
			return false
		},
		"filterRequired":    FilterRequired,
		"filterNotRequired": FilterNotRequired,
		"join":              strings.Join,
		"has": func(sl []string, str string) bool {
			for _, s := range sl {
				if s == str {
					return true
				}
			}
			return false
		},
		"toUpper":    strings.ToUpper,
		"toConstant": toConstant,
		"successfulResponse": func(responses openapi3.Responses) *openapi3.SchemaRef {
			for code, response := range responses {
				if code == "200" || code == "201" {
					return response.Value.Content["application/json"].Schema
				}
			}
			return nil
		},
		"requestSchema":      requestSchema,
		"operationContext":   operationContext,
		"sortedProperties":   sortedProperties,
		"requiredParameters": requiredParameters,
		"optionalParameters": optionalParameters,
		"clientReferences":   clientReferences,
		"additionalParameters": func() map[string]any {
			return config.AdditionalParameters
		},
	}
}

func requestSchema(operation *openapi3.Operation) *openapi3.SchemaRef {
	if operation.RequestBody == nil {
		return nil
	}

	if operation.RequestBody.Value.Content["application/json"] == nil {
		return nil
	}

	return operation.RequestBody.Value.Content["application/json"].Schema
}

type OperationWithPathAndMethod struct {
	Path      string
	Method    string
	Operation *openapi3.Operation

	RequestName       string
	RequestProperties []PropertyContext
}

func (o *OperationWithPathAndMethod) Imports() []string {

	imports := make([]string, 0)
	for _, prop := range o.RequestProperties {
		if prop.PropType.IsRef {
			imports = append(imports, prop.PropType.InnerType)
		}
	}
	return imports
}

type PropertyContext struct {
	Name     string
	PropType PropType
	Optional bool
}

// ByRequired implements sort.Interface for []PropertyContext based on
// the Optional field.
type ByRequired []PropertyContext

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (a ByRequired) Len() int           { return len(a) }
func (a ByRequired) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRequired) Less(i, j int) bool { return boolToInt(a[i].Optional) < boolToInt(a[j].Optional) }

type PropType struct {
	Ref       string
	IsRef     bool
	InnerType string
}

func NewPropType(ref string, isRef bool, innerType string) PropType {
	return PropType{Ref: ref, IsRef: isRef, InnerType: innerType}
}

func ConvertList(items *openapi3.SchemaRef) PropType {
	convertedType := ConvertType(items)
	ref := convertedType.Ref
	return NewPropType(fmt.Sprintf("List[%s]", ref), convertedType.IsRef, ref)
}

func ConvertDict(items *openapi3.SchemaRef) PropType {
	convertedType := ConvertType(items)
	ref := convertedType.Ref
	return NewPropType(fmt.Sprintf("Dict[str, %s]", ref), convertedType.IsRef, ref)
}

func ConvertFormat(format string) PropType {
	if format == "date-time" {
		return NewPropType("datetime", false, "")
	}
	return NewPropType("str", false, "")
}

// TODO: maybe factory pattern for other languages
func ConvertType(prop *openapi3.SchemaRef) PropType {
	t := prop.Value.Type

	if t == "" {
		return NewPropType("object", false, "")
	}
	if t == "string" {
		return ConvertFormat(prop.Value.Format)
	}
	if t == "integer" {
		return NewPropType("int", false, "")
	}
	if t == "number" {
		return NewPropType("float", false, "")
	}
	if t == "boolean" {
		return NewPropType("bool", false, "")
	}
	if t == "array" {
		return ConvertList(prop.Value.Items)
	}
	if t == "object" {
		if prop.Value.AdditionalProperties.Schema != nil {
			fmt.Println("additional properties: ", prop.Value.AdditionalProperties.Schema.Value.Type)

			return ConvertDict(prop.Value.AdditionalProperties.Schema)
		}
		return NewPropType(refToName(prop.Ref), true, refToName(prop.Ref))
	}
	if t == "null" {
		return NewPropType("None", false, "")
	}
	return NewPropType(t, false, "")

}

func operationContext(operation *openapi3.Operation, method, path string, typeContexts map[string]TypeContext) *OperationWithPathAndMethod {

	fmt.Println("operation: ", operation.OperationID)
	reqBody := operation.RequestBody
	var reqName string
	var typeContext TypeContext
	requestPropertiesSlice := make([]PropertyContext, 0)
	if reqBody != nil {
		reqName = refToName(reqBody.Value.Content["application/json"].Schema.Ref)
		fmt.Println("request body: ", reqName)
		typeContext = typeContexts[reqName]
		fmt.Println("name: ", typeContext.Name)
		properties := typeContext.Schema.Properties
		required := typeContext.Schema.Required

		for name, prop := range properties {
			isRequired := slices.Contains(required, name)

			requestPropertiesSlice = append(requestPropertiesSlice, PropertyContext{

				Name:     name,
				PropType: ConvertType(prop),
				Optional: !isRequired,
			},
			)
		}
	}
	//reorder requestProperties based on required, required parameters should be first in the list
	sort.Sort(ByRequired(requestPropertiesSlice))

	fmt.Println("request properties: ", requestPropertiesSlice)

	return &OperationWithPathAndMethod{
		Path:              path,
		Method:            method,
		Operation:         operation,
		RequestName:       reqName,
		RequestProperties: requestPropertiesSlice,
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

func refToName(ref string) string {
	return strings.TrimPrefix(ref, "#/components/schemas/")
}
