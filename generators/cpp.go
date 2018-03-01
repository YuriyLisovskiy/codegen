package generators

import (
	"../parser"
	"fmt"
	"strings"
)

var (
	cppClassFormat = "class %s%s\n{%s%s%s};"
	cppIndent      = getIndent(true, 4)
)

type CppGenerator struct{}

func (gen CppGenerator) Generate(pkg parser.Package) map[string]string {
	cppIndent = getIndent(!pkg.UseSpaces, 4)
	result := make(map[string]string)
	for _, class := range pkg.Classes {
		code := ""
		if class.Parent.Name != "" {
			code += "#include \"" + class.Parent.Name + ".h\"\n\n"
		}
		code += "using namespace std;\n\n"
		code += gen.generateClass(class) + "\n"
		code += "\n// Definition\n\n"
	//	code += "#include \"" + class.Name + ".h\"\n\nusing namespace std;\n\n"
		code += generateSourceFile(class, class.Name)
		result[class.Name] = code
	}
	return result
}

func (gen CppGenerator) generateClass(class parser.Class) string {
	parent, public, protected, private := "", "", "", ""
	if class.Parent.Name != "" {
		parent = " : " + class.Parent.Access + " " + class.Parent.Name
	}
	public = gen.generateSection("public", class)
	protected = gen.generateSection("protected", class)
	private = gen.generateSection("private", class)
	result := fmt.Sprintf(
		cppClassFormat,
		class.Name,
		parent,
		public,
		protected,
		private,
	)
	return result
}

func (CppGenerator) generateField(field parser.Field) string {
	result := cppIndent
	if field.Static {
		result += "static "
	}
	if field.Const {
		result += "const "
	}
	result += field.Type + " "
	result += field.Name
	if field.Default != "" {
		result += " = " + field.Default
	}
	result += ";"
	return result
}

func (CppGenerator) generateMethod(method parser.Method) string {
	result := ""
	if method.Static {
		result += "static "
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
		result += parameter.Type + parameter.Pass + " " + parameter.Name
		if i+1 < len(method.Parameters) {
			result += ", "
		}
	}
	result += ");"
	return result
}

func (gen CppGenerator) generateSection(access string, class parser.Class) string {
	result := ""
	for _, field := range class.Fields {
		if access == strings.ToLower(field.Access) {
			result += gen.generateField(field) + "\n"
		}
	}

	// TODO: generate constructors

	for _, method := range class.Methods {
		if access == strings.ToLower(method.Access) {
			result += shiftCode(gen.generateMethod(method), 1, cppIndent) + "\n"
		}
	}
	for _, class := range class.Classes {
		if access == strings.ToLower(class.Access) {
			result += shiftCode(gen.generateClass(class), 1, cppIndent) + "\n"
		}
	}
	if result != "" {
		result = "\n" + access + ":\n" + result
	}
	return result
}

func getReturnVal(returnType string) string {
	result := ""
	if strings.Contains(returnType, "*") {
		result = "nullptr"
	} else if strings.Contains(returnType, "int") {
		result = "0"
	} else if strings.Contains(returnType, "char") {
		result = "''"
	} else if strings.Contains(returnType, "double") {
		result = "0.0"
	} else {
		switch returnType {
		case "float":
			result = "0.0f"
		case "string":
			result = `""`
		default:
			result = ""
		}
	}
	return " " + result
}

func generateSourceFile(class parser.Class, access string) string {
	result := ""
	for _, method := range class.Methods {
		switch method.Return {
		case "":
			result += "void "
		default:
			result += method.Return + " "
		}
		result += access + "::" + method.Name + "("
		for i, parameter := range method.Parameters {
			if parameter.Const {
				result += "const "
			}
			result += parameter.Type + parameter.Pass + " " + parameter.Name
			if i+1 < len(method.Parameters) {
				result += ", "
			}
		}
		result += ")\n{\n"
		if method.Return != "" {
			result += cppIndent + "return" + getReturnVal(method.Return) + ";"
		}
		result += "\n}\n"
	}
	for _, cl := range class.Classes {
		result += generateSourceFile(cl, class.Name + "::" + cl.Name)
	}
	return result
}
