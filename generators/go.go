package generators

import (
	"fmt"
	"strings"
)

var (
	goClassFormat = "type %s struct {%s}%s%s"
	goIndent      = getIndent(true, 4)
)

type GoGenerator struct{}

func (gen GoGenerator) Generate(pkg Package) map[string]string {
	goIndent = getIndent(!pkg.UseSpaces, 4)
	result := make(map[string]string)
	for _, class := range pkg.Classes {
		code := "package " + class.Name + "\n\n"
		code += gen.generateClass(class)
		result[class.Name] = code
	}
	return result
}

func (gen GoGenerator) generateClass(class Class) string {
	fields, methods, classes := "", "", ""

	for _, field := range class.Fields {
		fields += gen.generateField(field) + "\n"
	}
	if fields != "" {
		fields = "\n" + fields
	}
	for _, method := range class.Methods {
		methods += "func (" + class.Name + ") " + gen.generateMethod(method) + "\n\n"
	}
	if methods != "" {
		methods = "\n\n" + methods
	}
	for _, innerClass := range class.Classes {
		classes += gen.generateClass(innerClass)
	}

	result := fmt.Sprintf(
		goClassFormat,
		class.Name,
		fields,
		methods,
		classes,
	)
	return result
}

func (GoGenerator) generateField(field Field) string {
	result := goIndent

	if field.Access == "public" {
		field.Name = strings.Title(field.Name)
	}

	result += field.Name + " " + field.Type

	return result
}

func (GoGenerator) generateMethod(method Method) string {
	result := ""

	if method.Access == "public" {
		method.Name = strings.Title(method.Name)
	}

	result += method.Name
	result += "("

	for i, parameter := range method.Parameters {
		result += parameter.Name + " " + parameter.Type
		if i+1 < len(method.Parameters) {
			result += ", "
		}
	}

	result += ")"

	if method.Return != "" {
		result += " " + method.Return
	}

	result += " {"

	if method.Return != "" {
		result += "\n" + goIndent + "return nil\n"
	}

	result += "}"

	return result
}

func (GoGenerator) generateGetSet(field []Field) string {
	return ""
}
