# openAPI Generator

## Usage

``` bash
go run . -i ../openapi/video-openapi.yaml -o ./go-generated -l go
```

## Writing Templates

Template language is go templates. [Documentation](https://golang.org/pkg/text/template/)

Template structure:

- `type.tmpl` - responsible for generating types/classes/structs and their fields. It should be called for each type in the spec from components section. Context is type definition from spec. It should call function to save generated type Name to type map. It is guaranteed that referenced types will be generated before this template is called. (TODO: what about recursive types or types that reference each other? Maybe we should generate typeNames separately and then generate types in any order)

- `request.tmpl` - responsible for generation of request type for each operation. It is called for each operation or /components/requests definition in the spec. Context is request definition.

- `operation.tmpl` - responsible for generating operations/methods. It should be called for each operation in the spec. Context is operation definition from spec and type Mappings. This template has abbility to register resulted function name in context (TODO: is it fine to have all operations in different files?)

- `client.tmpl` - generate interface/abstract class/etc for client. It should be called once for each client. Context is client definition from spec and type Mappings.

Templates are compiled into binary and distributed as a single binary. But for development purposes it is possible to build templates into separate files and use them from there using `--template-dir` flag.

Each target (might be language or language with specific framework) has own template directory. Each directory should have `config.yaml` file with configuration for template engine. It should contain information about parameters that are specific to this target, custom behaviour usage (e.g. generate dedicated types for each operation or not), type replacement (for example use time.Time for strings with format `date`).
<!-- We can use hardcoded values for such configuration because we have limited number of  languages to support. However I think keeping configuration and templates together is good option because templates may assume specific behaviour of engine. -->

TODO: describe config structure and how it affects template engine, describe template context, so it will be absolutely clear what is available in templates.

## API split. (Just an idea, not sure if it is good. It can reduce boilerplate in templates and share logic between targets)

We can generate several clients from one spec using one target, filtering the spec using extensions or tags. So we can have one template and call entire process for backend/frontend integration or split client into 2: one for channels api, one for messages api.

Filter values should be specified in config.yaml or/and in command line arguments. For example:

in config.yaml:

```yaml
allowed_integration_type: [backend] # number of types are limited by engine
```

this line specifies that this target templates can be used for backend integration only

in command line:

```sh
openapi-gen -i <input_spec> -o <output-dir> -t <template-dir> --integration-type backend
```

Running this command will generate only backend integration client. Passing `frontend` as integration type will result in error.

## Feature list

- [ ] Generate simple types
- [ ] Generate oneOf types for events.
- [ ] Generate request types
- [ ] Generate client type
- [ ] Generate methods for client. Methods should send requests and return responses.

| Feature\Target        | Python    | Swift | Kotlin | Typescript | Go |
|:----------------------|:---------:|:-----:|:------:|:----------:|:--:|
| Simple types          |Data Class + dataclass json|       |        |            | struct with json tags |
| OneOf types for events| ?? |       |        |            |  Interface with private method and limited structs which has this method.  |
| Request types         | Data Class       |       |        |            |struct |
| Client type           |  class with method sending request, can be prewritten without generation       |       |        |            |  struct and generic method to send request, handle errors, parse response  |
| Client methods        |   just a generated functions      |       |        |            |  methods using 1 method to send request  |
