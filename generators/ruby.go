package generators

import (
	"fmt"
	"strings"
)

var (
	rubyClassFormat = "class %s %s\n%s%s%send"
	rubyIndent      = getIndent(true, 4)
)

type RubyGenerator struct{}

func (gen RubyGenerator) Generate(pkg Package) map[string]string {
	rubyIndent = getIndent(!pkg.UseSpaces, 4)
	result := make(map[string]string)
	for _, class := range pkg.Classes {
		code := gen.generateClass(class) + "\n"
		result[class.Name] = code
	}
	return result
}

func (gen RubyGenerator) generateClass(class Class) string {
	fields, inherits, methods, classes := "", "", "", ""

	if class.Parent.Name != "" {
		inherits = "< " + class.Parent.Name
	}

	if len(class.Fields) > 0 {
		fields = shiftCode(gen.generateInit(class), 1, rubyIndent) + "\n"
	}
	for _, method := range class.Methods {
		methods += "\n" + shiftCode(gen.generateMethod(method), 1, rubyIndent) + "\n"
	}
	for _, innerClass := range class.Classes {
		classes += "\n" + shiftCode(gen.generateClass(innerClass), 1, rubyIndent) + "\n"
	}

	result := fmt.Sprintf(
		rubyClassFormat,
		class.Name,
		inherits,
		fields,
		methods,
		classes,
	)
	return result
}

func (RubyGenerator) generateField(field Field) string {
	result := rubyIndent

	if field.Access == "public" {
		field.Name = strings.Title(field.Name)
	}

	result += field.Name + " " + field.Type

	return result
}

func (gen RubyGenerator) generateMethod(method Method) string {
	return gen.generateMethodWithBody(method, "")
}

func (RubyGenerator) generateMethodWithBody(method Method, body string) string {
	result := "def "

	if method.Access == "private" {
		result = "private " + result
	}
	if method.Static {
		result += "self."
	}
	result += method.Name

	if len(method.Parameters) > 0 {
		result += "("
		for i, parameter := range method.Parameters {
			result += parameter.Name
			if parameter.Default != "" {
				result += "=" + parameter.Default
			}
			if i+1 < len(method.Parameters) {
				result += ", "
			}
		}
		result += ")"
	}

	//if method.Return != "" {
	//	result += " " + method.Return
	//}

	result += "\n"
	result += shiftCode(body, 1, rubyIndent)
	//if method.Return != "" {
	//	result += "\n" + rubyIndent + "return nil\n"
	//}

	result += "\nend"

	return result
}

func (gen RubyGenerator) generateInit(class Class) string {
	result, body := "", ""
	init := Method{
		Name:       "initialize",
		Parameters: []Parameter{},
	}
	for i, field := range class.Fields {
		init.Parameters = append(init.Parameters, Parameter{
			Name:    field.Name,
			Default: field.Default,
		})
		body += "@" + field.Name + " = " + field.Name
		if i+1 < len(class.Fields) {
			body += "\n"
		}
	}
	result += gen.generateMethodWithBody(init, body)

	return result
}
