package generators

import (
	"../parser"
	"fmt"
	"strings"
)

var (
	pythonClassFormat = "class %s%s:\n%s%s%s"
	pythonIndent      = getIndent(true, 4)
)

type PythonGenerator struct{}

func (gen PythonGenerator) Generate(pkg parser.Package) string {
	pythonIndent = getIndent(!pkg.UseSpaces, 4)
	result := ""
	for _, class := range pkg.Classes {
		result += parser.DELIM_START + gen.generateClass(class) + "\n" + parser.DELIM_END
	}
	return result
}

func (gen PythonGenerator) generateClass(class parser.Class) string {
	fields, inherits, methods, classes := "", "", "", ""

	if class.Parent.Name != "" {
		inherits = "(" + class.Parent.Name + ")"
	}

	if len(class.Fields) > 0 {
		fields = shiftCode(gen.generateInit(class), 1, pythonIndent)
	}
	for _, method := range class.Methods {
		methods += "\n" + shiftCode(gen.generateMethod(method), 1, pythonIndent) + "\n"
	}
	for _, innerClass := range class.Classes {
		classes += "\n" + shiftCode(gen.generateClass(innerClass), 1, pythonIndent)
	}

	result := fmt.Sprintf(
		pythonClassFormat,
		class.Name,
		inherits,
		fields,
		methods,
		classes,
	)

	if result[len(result)-2:] == ":\n" {
		result += pythonIndent + "pass"
	}
	return result
}

func (PythonGenerator) generateField(field parser.Field) string {
	result := goIndent

	if field.Access == "public" {
		field.Name = strings.Title(field.Name)
	}

	result += field.Name + " " + field.Type

	return result
}

func (gen PythonGenerator) generateMethod(method parser.Method) string {
	return gen.generateMethodWithBody(method, "pass")
}

func (PythonGenerator) generateMethodWithBody(method parser.Method, body string) string {
	result := "def "

	switch method.Access {
	case "private":
		method.Name = "__" + method.Name
	case "protected":
		method.Name = "_" + method.Name
	}

	result += method.Name
	result += "("

	if method.Static {
		result = "@staticmethod\n" + result
	} else {
		result += "self"
		if len(method.Parameters) > 0 {
			result += ", "
		}
	}

	for i, parameter := range method.Parameters {
		result += parameter.Name
		if parameter.Default != "" {
			result += "=" + parameter.Default
		}
		if i+1 < len(method.Parameters) {
			result += ", "
		}
	}
	result += "):"

	//if method.Return != "" {
	//	result += " " + method.Return
	//}

	result += "\n"
	result += shiftCode(body, 1, pythonIndent)
	//if method.Return != "" {
	//	result += "\n" + goIndent + "return nil\n"
	//}
	return result
}

func (gen PythonGenerator) generateInit(class parser.Class) string {
	result, body := "", ""
	init := parser.Method{
		Name:       "__init__",
		Parameters: []parser.Parameter{},
	}
	for _, field := range class.Fields {
		init.Parameters = append(init.Parameters, parser.Parameter{
			Name:    field.Name,
			Default: field.Default,
		})
		body += "self." + field.Name + " = " + field.Name + "\n"
	}
	result += gen.generateMethodWithBody(init, body)

	return result
}
