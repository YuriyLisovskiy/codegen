package generators

import (
	"errors"
	"strings"
)

type Generator interface {
	Generate(class Package) map[string]string
	generateField(field Field) string
	generateMethod(method Method) string
	generateClass(class Class) string
}

type Language struct {
	Generator Generator
	Extension string
	Comment   string
}

var (
	ExamplePkg = Package{
		Name:      "main",
		UseSpaces: false,
		Classes: []Class{
			{
				Name: "Fruit",
			},
			{
				Name: "Apple",
				Parent: Parent{
					Name:   "Fruit",
					Access: "public",
				},
				Fields: []Field{
					{
						Access:  "public",
						Type:    "string",
						Name:    "colour",
						Default: `"red"`,
					},
					{
						Access:  "public",
						Type:    "string",
						Static:  true,
						Name:    "sort",
						Default: `"Golden"`,
					},
					{
						Access:  "private",
						Type:    "int",
						Name:    "size",
						Default: "1",
					},
				},
				Methods: []Method{
					{
						Access: "private",
						Name:   "print",
						Parameters: []Parameter{
							{
								Pass:  "&",
								Name:  "colour",
								Type:  "string",
								Const: true,
							},
						},
					},
					{
						Access: "protected",
						Return: "int",
						Static: true,
						Name:   "getSize",
					},
					{
						Access: "public",
						Return: "string",
						Name:   "getColor",
						Const:  true,
					},
				},
			},
			{
				Access: "private",
				Name:   "Seed",
				Fields: []Field{
					{
						Access: "public",
						Type:   "int",
						Name:   "size",
					},
				},
				Methods: []Method{
					{
						Static: true,
						Access: "public",
						Return: "int",
						Name:   "transform",
						Const:  true,
					},
				},
			},
		},
	}
	Generators = map[string]Language{
		"java":   {&JavaGenerator{}, "java", "/* %s */"},
		"go":     {&GoGenerator{}, "go", "/* %s */"},
		"ruby":   {&RubyGenerator{}, "rb", "# %s\n"},
		"cpp":    {&CppGenerator{}, "h", "/* %s */"},
		"python": {&PythonGenerator{}, "py", "# %s\n"},
		"js_es6": {&ES6Generator{}, "js", "/* %s */"},
		"csharp": {&CSharpGenerator{}, "cs", "/* %s */"},
		"xml":    {&XmlGenerator{}, "xml", "<!-- %s -->"},
		"json":   {&JsonGenerator{}, "json", ""},
		"yaml":   {&YamlGenerator{}, "yml", "# %s\n"},
	}
)

func getIndent(tabs bool, tabStop int) string {
	if tabs {
		return "\t"
	} else {
		return strings.Repeat(" ", tabStop)
	}
}

func shiftCode(code string, num int, indent string) string {
	indent = strings.Repeat(indent, num)
	return indent + strings.Replace(code, "\n", "\n"+indent, -1)
}

func NormalizeLang(lang string) string {
	if lang == "js" {
		lang = "js_es6"
	} else if lang == "c#" || lang == "cs" {
		lang = "csharp"
	} else if lang == "yml" {
		lang = "yaml"
	} else if lang == "c++" {
		lang = "cpp"
	}
	return lang
}

func GetGenerator(name string) (Generator, error) {
	name = NormalizeLang(name)
	gen := Generators[name]
	if gen.Generator == nil {
		return nil, errors.New("this generator doesn't exist")
	}
	return gen.Generator, nil
}
