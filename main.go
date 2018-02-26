package main

import (
	"./generators"
	"./parser"
	"fmt"
	"flag"
	"errors"
)

var (
	lang  = "csharp"
	class = parser.Class{
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
		Classes: []parser.Class{
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


func getArgs() (string, string, string, bool) {
	langPtr := flag.String("l", "", "language")
	xmlPtr := flag.String("f", "", "file")
	xmlUrlPtr := flag.String("u", "", "file url")
	spacesPtr := flag.Bool("s", false, "use spaces")
	flag.Parse()
	return *langPtr, *xmlPtr, *xmlUrlPtr, *spacesPtr
}

func parseArgs(lang, xml, url string, spaces bool) (string, []byte, error) {
	
	file := []byte("")
	if lang == "" {
		return "", file, errors.New("specify language (-l) flag")
	}
	if url == "" && xml == "" {
		return "", file, errors.New("specify file path (-f) or url path (-u) flag")
	}
	if xml != "" && url != "" {
		return "", file, errors.New("do not use both -f and -u flags at the same time")
	}
	if xml != "" {
		file, err := parser.Read(xml)
		if err != nil {
			return "", file, err
		}
		return lang, file, nil
	} else if url != "" {
		file, err := parser.Download(url)
		if err != nil {
			return "", file, err
		}
		return lang, file, nil
	}
	return lang, file, nil
}

func main() {
	
//	language, file, err := parseArgs(getArgs())
//	if err != nil {
//		panic(err)
//	}
	generator, err := generators.GetGenerator(lang)
	if err != nil {
		panic(err)
	}
//	fmt.Println(generator.Generate(parser.Parse(file)))
	fmt.Println(generator.Generate(class))
}
