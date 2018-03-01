package main

import (
	"fmt"
	"errors"
	"./parser"
	"./generators"
)

var (
	lang  = "csharp"
	v = parser.Package{
		Classes: []parser.Class{
			{
				Name: "Apple",
				Parent: parser.Parent{
					Name:   "Fruit",
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
						Access:  "public",
						Type:    "string",
						Static:  true,
						Name:    "sort",
						Default: `"Golden"`,
					},
					{
						Access:  "private",
						Type:    "int",
						Name:    "size",
						Default: "1",
					},
				},
				Methods: []parser.Method{
					{
						Access: "private",
						Name:   "print",
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
						Static: true,
						Name:   "getSize",
					},
					{
						Access: "public",
						Return: "string",
						Name:   "getColor",
						Const:  true,
					},
				},
			},
			{
				Access: "private",
				Name:   "Seed",
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
						Name:   "transform",
						Const:  true,
					},
				},
			},
		},
	}
)


func execute() error {
	language, fileName, url, useSpaces := parser.GetArgs()
	err := parser.ValidateArgs(language, fileName, url)
	if err != nil {
		return err
	}
	var byteContext []byte
	if url != "" {
		byteContext, err = parser.Download(url)
		if err != nil {
			return err
		}
		fileName = url
	} else {
		byteContext, err = parser.Read(fileName)
		if err != nil {
			return err
		}
	}
	generator, err := generators.GetGenerator(language)
	if err != nil {
		return nil
	}
	fileFormat, err := parser.GetFileFormat(fileName)
	if err != nil {
		return err
	}
	var object parser.Package
	switch fileFormat {
	case "xml":
		object = parser.ParseXml(byteContext)
	case "json":
		object = parser.ParseJson(byteContext)
	case "yml":
		object = parser.ParseYaml(byteContext)
	default:
		return errors.New(fmt.Sprintf("Invalid format of '%s' file.", fileName))
	}
	object.UseSpaces = useSpaces
	fileContextMap := generator.Generate(object)
	ext := parser.GetExtension(language)
	for key, val := range fileContextMap {
		err = parser.Write(key + ext, val)
		if err != nil {
			return err
		}
	}
	if len(object.Classes) == 0 && len(object.Classes) == 0 && len(object.Classes) == 0 {
		fmt.Println(fmt.Sprintf("There is nothing to generate, file '%s' is empty or incorrect.", fileName))
	} else {
		fmt.Println("Generated successfully.")
	}
	return nil
}

func main() {
//	generator, err := generators.GetGenerator(lang)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(generator.Generate(object))

	err := execute()
	if err != nil {
		panic(err)
	}

}
