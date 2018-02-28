package generators

import (
	"../parser"
	"fmt"
	"strings"
)

var (
	es6ClassFormat = "class %s %s{%s%s}%s"
	es6Indent      = getIndent(true, 4)
)

type ES6Generator struct{}

func (gen ES6Generator) Generate(pkg parser.Package) string {
	es6Indent = getIndent(!pkg.UseSpaces, 4)
	result := ""
	for _, class := range pkg.Classes {
		result += parser.DELIM_START + gen.generateClass(class) + "\n" + parser.DELIM_END
	}
	return result
}

func (gen ES6Generator) generateClass(class parser.Class) string {
	fields, inherits, methods, classes := "", "", "", ""

	if class.Parent.Name != "" {
		inherits = "extends " + class.Parent.Name + " "
	}

	if len(class.Fields) > 0 {
		fields = "\n" + shiftCode(gen.generateInit(class), 1, es6Indent) + "\n"
	}

	for _, method := range class.Methods {
		methods += "\n" + shiftCode(gen.generateMethod(method), 1, es6Indent) + "\n"
	}
	for _, innerClass := range class.Classes {
		classes += "\n" + class.Name + "." + innerClass.Name + " = " + gen.generateClass(innerClass)
	}

	result := fmt.Sprintf(
		es6ClassFormat,
		class.Name,
		inherits,
		fields,
		methods,
		classes,
	)

	for _, field := range class.Fields {
		if field.Static {
			result += "\n" + class.Name + "." + field.Name + " = "
			if field.Default != "" {
				result += field.Default
			} else {
				result += "null"
			}
		}
	}

	return result
}

func (ES6Generator) generateField(field parser.Field) string {
	result := es6Indent

	if field.Access == "public" {
		field.Name = strings.Title(field.Name)
	}

	result += field.Name + " " + field.Type

	return result
}

func (gen ES6Generator) generateMethod(method parser.Method) string {
	return gen.generateMethodWithBody(method, "")
}

func (ES6Generator) generateMethodWithBody(method parser.Method, body string) string {
	result := ""

	//if method.Access == "private" {
	//	result = "private " + result
	//}
	if method.Static {
		result += "static "
	}
	result += method.Name

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
	result += ") {"

	//if method.Return != "" {
	//	result += " " + method.Return
	//}

	result += ""
	if body != "" {
		result += "\n" + shiftCode(body, 1, es6Indent) + "\n"
	}
	//if method.Return != "" {
	//	result += "\n" + es6Indent + "return nil\n"
	//}

	result += "}"

	return result
}

func (gen ES6Generator) generateInit(class parser.Class) string {
	result, body := "", ""
	init := parser.Method{
		Name:       "constructor",
		Parameters: []parser.Parameter{},
	}
	previousIsStatic := false
	for i, field := range class.Fields {
		if field.Static {
			previousIsStatic = true
			continue
		}
		init.Parameters = append(init.Parameters, parser.Parameter{
			Name:    field.Name,
			Default: field.Default,
		})
		body += "this." + field.Name + " = " + field.Name
		if i+1 < len(class.Fields) && !previousIsStatic {
			body += "\n"
		}
		previousIsStatic = false
	}
	if previousIsStatic && len(body) > len(es6Indent) {
		body = body[:len(body)-len(es6Indent)+1]
	}
	if class.Parent.Name != "" {
		body = "super()\n" + body
	}

	result += gen.generateMethodWithBody(init, body)

	return result
}
