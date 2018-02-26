package generators

import (
	"../parser"
	"errors"
	"strings"
)

type Generator interface {
	Generate(class parser.Class) string
	generateField(field parser.Field) string
	generateMethod(method parser.Method) string
	generateClass(class parser.Class) string
}

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

func GetGenerator(name string) (Generator, error) {
	switch name {
	case "java":
		return &JavaGenerator{}, nil
	case "go":
		return &GoGenerator{}, nil
	case "ruby":
		return &RubyGenerator{}, nil
	case "cpp":
		return &CppGenerator{}, nil
	case "python":
		return &PythonGenerator{}, nil
	case "js_es6":
		return &ES6Generator{}, nil
	case "csharp":
		return &CSharpGenerator{}, nil
	}
	return nil, errors.New("this generator doesn't exist")
}
