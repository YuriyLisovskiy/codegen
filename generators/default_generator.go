package generators

import (
	"errors"
//	"github.com/YuriyLisovskiy/codegen/parser"
	"strings"
	"../parser"
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
	case "cpp":
		return &CppGenerator{}, nil
	}
	return nil, errors.New("this generator doesn't exist")
}
