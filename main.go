package main

import (
	"fmt"
	"github.com/YuriyLisovskiy/codegen/generators"
	"github.com/YuriyLisovskiy/codegen/parser"
)

var (
	lang  = "go"
	class = parser.Class{
		Name: "Apple",
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
						Name:  "colour",
						Type:  "string",
						Const: true,
					},
				},
			},
			{
				Access: "private",
				Return: "int",
				Name:   "getSize",
			},
		},
		Classes: []parser.Class{
			{
				Name: "Seed",
				Fields: []parser.Field{
					{
						Access: "public",
						Type:   "int",
						Name:   "size",
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
