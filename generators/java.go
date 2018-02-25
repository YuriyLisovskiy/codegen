package generators

import (
	"fmt"
//	"github.com/YuriyLisovskiy/codegen/parser"
	"../parser"
)

var (
	javaClassFormat = "class %s {%s%s%s}"
	javaIndent      = getIndent(true, 4)
)

type JavaGenerator struct{}

/*
The class must be validated before using this function
*/
func (gen JavaGenerator) Generate(class parser.Class) string {
	return gen.generateClass(class)
}

func (gen JavaGenerator) generateClass(class parser.Class) string {
	fields, methods, classes := "", "", ""

	for _, field := range class.Fields {
		fields += gen.generateField(field) + "\n"
	}
	if fields != "" {
		fields = "\n" + fields
	}
	for _, method := range class.Methods {
		methods += shiftCode(gen.generateMethod(method), 1, javaIndent) + "\n\n"
	}
	if methods != "" {
		methods = "\n" + methods
	}
	for _, innerClass := range class.Classes {
		classes += shiftCode(gen.generateClass(innerClass), 1, javaIndent)
	}
	if classes != "" {
		classes = classes + "\n"
	}

	result := fmt.Sprintf(
		javaClassFormat,
		class.Name,
		fields,
		methods,
		classes,
	)
	return result
}

func (JavaGenerator) generateField(field parser.Field) string {
	result := javaIndent
	if field.Access == "" || field.Access == "default" {
		result += "private "
	} else {
		result += field.Access + " "
	}
	if field.Const {
		result += "const "
	}
	switch field.Type {
	case "string":
		result += "String "
	default:
		result += field.Type + " "
	}

	result += field.Name

	if field.Default != "" {
		result += " = " + field.Default
	}
	result += ";"
	return result
}

func (JavaGenerator) generateMethod(method parser.Method) string {
	result := ""
	if method.Access == "" || method.Access == "default" {
		result += "private "
	} else {
		result += method.Access + " "
	}
	if method.Return == "string" {
		method.Return = "String"
	}
	switch method.Return {
	case "":
		result += "void "
	default:
		result += method.Return + " "
	}

	result += method.Name
	result += "("

	for i, parameter := range method.Parameters {
		if parameter.Const {
			result += "const "
		}
		if parameter.Type == "string" {
			parameter.Type = "String"
		}
		result += parameter.Type + " " + parameter.Name
		if i+1 < len(method.Parameters) {
			result += ", "
		}
	}

	result += ") {"

	if method.Return != "" {
		result += "\n" + javaIndent + "return new " + method.Return + "();\n"
	}

	result += "}"

	return result
}
