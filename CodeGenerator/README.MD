# CodeGenerator
--
#### Install

    go get "github.com/ivan-kostko/GoLibs/CodeGenerator"

    // To define any other location or/and template file extention for templates - use build flags like the following:
    go get -ldflags "-X main.templatesMainFolder=YourCustomTemplateLocation -X main.templatesSubFolder=" "github.com/ivan-kostko/GoLibs/CodeGenerator"

#### Description

    The project CodeGenerator represents //go:generate tool to generate _.go files from templates
    It uses standard text/template notation.
    To prevent template differencies it is loking for templates at get project forlder subfolder (default=CodeGeneratorTemplates) files with predefined extention(default=tmpl)

#### Example

    //go:generate CodeGenerator -pointer -template=container -type=MyType

    It will generate in currently active folder file mytype_container.go with package named as current folder name.
