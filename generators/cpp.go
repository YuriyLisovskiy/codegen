package generators

import (
	"fmt"
	//	"github.com/YuriyLisovskiy/codegen/parser"
	"../parser"
	"strings"
)

var (
	cppClassFormat = "class %s%s\n{%s%s%s}"
	cppIndent      = getIndent(true, 4)
)

type CppGenerator struct{}


func (gen CppGenerator) Generate(class parser.Class) string {
	return gen.generateClass(class)
}

func (gen CppGenerator) generateClass(class parser.Class) string {
	parents, public, protected, private := "", "", "", ""
	for i, parent := range class.Parents {
		parents += parent.Access + " " + parent.Name
		if i + 1 < len(class.Parents) {
			parents += ", "
		}
	}
	if parents != "" {
		parents = " : " + parents
	}
	public = gen.generateSection("public", class)
	protected = gen.generateSection("protected", class)
	private = gen.generateSection("private", class)
	result := fmt.Sprintf(
		cppClassFormat,
		class.Name,
		parents,
		public,
		protected,
		private,
	)
	return result
}

func (CppGenerator) generateField(field parser.Field) string {
	result := javaIndent
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
	result += ")\n{"
	if method.Return != "" {
		result += "\n" + javaIndent + "return " + method.Return + ";\n"
	}
	result += "}"
	return result
}

func (gen CppGenerator) generateSection(access string, class parser.Class) string {
	result := ""
	for _, field := range class.Fields {
		if access == strings.ToLower(field.Access) {
			result += shiftCode(gen.generateField(field), 1, cppIndent) + "\n"
		}
	}
	
	// TODO: generate constructors
	
	for _, method := range class.Methods {
		if access == strings.ToLower(method.Access) {
			result += shiftCode(gen.generateMethod(method), 1, cppIndent) + "\n\n"
		}
	}
	for _, class := range class.Classes {
		if access == strings.ToLower(class.Access) {
			result += shiftCode(gen.generateClass(class), 1, cppIndent) + "\n\n"
		}
	}
	if result != "" {
		result = access + ":\n" + result + "\n"
	}
	return result
}
