package generators

import (
	"fmt"
	"strings"
)

var (
	cppClassFormat = "class %s%s\n{%s%s%s};"
	cppIndent      = getIndent(true, 4)
)

type CppGenerator struct{}

func (gen CppGenerator) Generate(pkg Package) map[string]string {
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

func (gen CppGenerator) generateClass(class Class) string {
	parent, public, protected, private := "", "", "", ""
	if class.Parent.Name != "" {
		parent = " : " + class.Parent.Access + " " + class.Parent.Name
	}
	public = gen.generateSection("public", class)
	if public != "" {
		public = "\n" + public
	}
	getSet := gen.generateGetSet(class.Fields)
	if getSet != "" {
		if public == "" {
			public += "\npublic:\n"
		}
		public += getSet
	}
	protected = gen.generateSection("protected", class)
	if protected != "" {
		protected = "\n" + protected
	}
	private = gen.generateSection("private", class)
	if private != "" {
		private = "\n" + private
	}
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

func (CppGenerator) generateField(field Field) string {
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

func (CppGenerator) generateMethod(method Method) string {
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

func (CppGenerator) generateGetSet(fields []Field) string {
	result := ""
	for _, field := range fields {
		if field.Getter {
			result += cppIndent + generateGet(field, "") + "\n"
		}
		if field.Setter {
			result += cppIndent + generateSet(field, "") + "\n"
		}
	}
	return result
}

func (gen CppGenerator) generateSection(access string, class Class) string {
	result := ""
	for _, field := range class.Fields {
		if access == strings.ToLower(field.Access) {
			result += gen.generateField(field) + "\n"
		}
	}
	if result != "" {
		result += "\n"
	}
	if  access == "public" && len(class.Fields) > 0 {
		result += shiftCode(generateConstructors(class, ""), 1, cppIndent) + "\n"
	}
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
		result = access + ":\n" + result
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

func generateSourceFile(class Class, access string) string {
	result := generateConstructors(class, access)
	if result != "" {
		result = "\n" + result + "\n"
	}
	for _, method := range class.Methods {
		result += "\n"
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
	for _, field := range class.Fields {
		if field.Getter {
			result += "\n" + generateGet(field, access) + "\n"
		}
		if field.Setter {
			result += "\n" + generateSet(field, access) + "\n"
		}
	}
	for _, cl := range class.Classes {
		result += "\n" + generateSourceFile(cl, class.Name+"::"+cl.Name) + "\n"
	}
	return result
}

func generateGet(field Field, access string) string {
	result := ""
	if access != "" {
		result = field.Type + " " + access + "::get" + strings.Title(field.Name) + "()\n{\n" +
			cppIndent + "return " + field.Name + ";\n}"
	} else {
		result = field.Type + " get" + strings.Title(field.Name) + "();"
	}
	return result
}

func generateSet(field Field, access string) string {
	result := ""
	if access != "" {
		result = "void " + access + "::set" + strings.Title(field.Name) + "(" + field.Type +
			" newValue)\n{\n" + cppIndent + "this->" + field.Name + " = newValue;\n}"
	} else {
		result = "void set" + strings.Title(field.Name) + "(" + field.Type + " newValue);"
	}
	return result
}

func generateConstructors(class Class, access string) string {
	params := ""
	for i, field := range class.Fields {
		params += field.Type + " " + field.Name
		if i+1 < len(class.Fields) {
			params += ", "
		}
	}
	result := ""
	if access != "" {
		if params != "" {
			result += "\n" + access + "::" + class.Name + "(" + params + ")\n{\n"
			for _, field := range class.Fields {
				result += cppIndent + "this->" + field.Name + " = " + field.Name + ";\n"
			}
			result += "}"
		}
	} else {
		result = class.Name + "() = default;"
		if params != "" {
			result += "\n" + class.Name + "(" + params + ");"
		}
	}
	return result
}
