package generators

import (
	"fmt"
	"strings"
)

var (
	cSharpClassFormat = "class %s %s{%s%s%s}"
	cSharpIndent      = getIndent(true, 4)
)

type CSharpGenerator struct{}

/*
The class must be validated before using this function
*/
func (gen CSharpGenerator) Generate(pkg Package) map[string]string {
	cSharpIndent = getIndent(!pkg.UseSpaces, 4)
	result := make(map[string]string)
	for _, class := range pkg.Classes {
		result[class.Name] = gen.generateClass(class)
	}
	return result
}

func (gen CSharpGenerator) generateClass(class Class) string {
	fields, inherits, methods, classes := "", "", "", ""

	if class.Parent.Name != "" {
		inherits = ": " + class.Parent.Name + " "
	}
	for _, field := range class.Fields {
		fields += gen.generateField(field) + "\n"
	}
	if fields != "" {
		fields = "\n" + fields
	}
	for _, method := range class.Methods {
		methods += "\n" + shiftCode(gen.generateMethod(method), 1, cSharpIndent) + "\n"
	}
	if len(class.Fields) > 0 {
		methods += shiftCode(gen.generateGetSet(class.Fields), 1, cSharpIndent)
	}
	for _, innerClass := range class.Classes {
		classes += "\n" + shiftCode(gen.generateClass(innerClass), 1, cSharpIndent) + "\n"
	}
	if classes != "" {
		classes += "\n"
	} else if methods != "" {
		methods += "\n"
	} else if fields != "" {
		fields += "\n"
	}
	result := fmt.Sprintf(
		cSharpClassFormat,
		class.Name,
		inherits,
		fields,
		methods,
		classes,
	)
	return result
}

func (CSharpGenerator) generateField(field Field) string {
	result := javaIndent
	if field.Access == "" || field.Access == "default" {
		result += "private "
	} else {
		result += field.Access + " "
	}
	if field.Const {
		result += "const "
	}
	if field.Static {
		result += "static "
	}
	switch field.Type {
	case "string":
		result += "System.String "
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

func (CSharpGenerator) generateMethod(method Method) string {
	result := ""
	if method.Access == "" || method.Access == "default" {
		result += "private "
	} else {
		result += method.Access + " "
	}
	if method.Static {
		result += "static "
	}
	if method.Return == "string" {
		method.Return = "System.String"
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
			parameter.Type = "System.String"
		}
		result += parameter.Type + " " + parameter.Name
		if i+1 < len(method.Parameters) {
			result += ", "
		}
	}

	result += ") {"

	if method.Return != "" {
		switch method.Return {
		case "string":
			result += "\n" + javaIndent + "return new System.String();\n"
		default:
			result += "\n" + javaIndent + "return new " + method.Return + "();\n"
		}
	}

	result += "}"

	return result
}

func (gen CSharpGenerator) generateGetSet(fields []Field) string {
	result := ""
	for _, field := range fields {
		if field.Getter {
			result += "\npublic " + field.Type + " get" + strings.Title(field.Name) + "() {\n" +
				cSharpIndent + "return " + field.Name + ";\n}\n"
		}
		if field.Setter {
			result += "\npublic void set" + strings.Title(field.Name) + "(" + field.Type + " new" +
				strings.Title(field.Name) + ") {\n" + cSharpIndent + field.Name + " = new" +
					strings.Title(field.Name) + ";\n}\n"
		}
	}
	return result
}
