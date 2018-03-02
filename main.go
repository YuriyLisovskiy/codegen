package main

import (
	"fmt"
	"github.com/YuriyLisovskiy/codegen/generators"
)

var (
	defaultLang = "xml"
	defaultPkg  = generators.ExamplePkg
)

func main() {
	//	generator, err := generators.GetGenerator(lang)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(generator.Generate(object))

	err := execute()

	if err != nil {
		fmt.Println(err)
	}
}
