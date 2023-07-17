# openAPI Generator

## Usage

``` bash
opeanapi-gen -i <input_spec> -o <output> (-t <template-dir> or target language) -integration <client|server>
```

- intergation parameter hides/shows parts of the spec that are not relevant for client/server code generation to template engine. If we pass client we hide some paramaeters and operations. (do we have client only parameters, operations?)

<!-- something about client/sever code so we can add/remove parts related only to one of them -->
<!-- do we need some context specific information? like package name, etc. should be as low as possible -->

## Writing Templates

<!-- 
    - template language (mustache and handlebars)
    - templates structure, so we don't put everything in one file
    - template variables and functions, so we can reuse common logic like filtering part of the spec, pass context between templates, etc.
    - common practices we enforce
    - calling other templates
    - testing and debugging templates?
    - how to pass docs from spec to code
 -->

 <!-- current template structure
 
  -->

Template structure:

- api - template called once for spec file, params - spec
- operation - template called for each operation, params - operation, models
- model - template called for each declared type in components
- other files can be called from templates, but they are not called directly

<!-- We can embed templates in binary or load them from disk. If we embed them we'll be able to install go get to install that tool without any additional steps. If we load them from disk we'll be able to modify them without recompiling the tool. Anyway we will store templates in same repo as generator. -->

Other structure:

- put 1 file per language, which called with entire spec, and then it calls other templates, seems like openapi-generator does it this way, bot it is not clear what kind of file layet it produces
- api.mustache - root template.

Available variables and functions:

- spec - entire spec
- package - package name ?? that param probably language specific and should be passed from cli. So we need to specify required params for each template

<!-- questions to think about: what kind of code should be generated for allOf types, so can we compile it as it one struct or better to let templates know, what go -> openAPI generator doing with composition? now we don't have allOf in specs, so suppose generator handles it and ecpand it to simple structure -->

### Testing generated code??
### Debugging templates??