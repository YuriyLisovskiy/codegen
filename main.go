package main

import (
	"fmt"
//	"github.com/YuriyLisovskiy/codegen/generators"
//	"github.com/YuriyLisovskiy/codegen/parser"
	"./parser"
	"./generators"
)

var (
	lang  = "cpp"
	class = parser.Class{
		Name: "Apple",
		Parent: parser.Parent{
			Name: "Fruit",
			Access: "public",
		},
		Fields: []parser.Field{
			{
				Access:  "public",
				Type:    "string",
				Name:    "colour",
				Default: `"red"`,
			},
			{
				Access:  "private",
				Type:    "int",
				Name:    "size",
				Default: `1`,
			},
		},
		Methods: []parser.Method{
			{
				Access: "private",
				Name:   "printColour",
				Parameters: []parser.Parameter{
					{
						Pass:  "&",
						Name:  "colour",
						Type:  "string",
						Const: true,
					},
				},
			},
			{
				Access: "protected",
				Return: "int",
				Name:   "getSize",
			},
			{
				Access: "public",
				Return: "string",
				Name:   "getColor",
				Const:  true,
			},
		},
		Classes: []parser.Class{
			{
				Access: "private",
				Name: "Seed",
				Fields: []parser.Field{
					{
						Access: "public",
						Type:   "int",
						Name:   "size",
					},
				},
				Methods: []parser.Method{
					{
						Static: true,
						Access: "public",
						Return: "int",
						Name:   "getSize",
						Const:  true,
					},
				},
			},
		},
	}
)

/*
	Parse command arguments using "flag" library
	read more here: https://gobyexample.com/command-line-flags
*/
func main() {

	generator, err := generators.GetGenerator(lang)
	if err != nil {
		panic(err)
	}
	fmt.Println(generator.Generate(class))
}
