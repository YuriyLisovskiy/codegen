package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"

	"flag"
	"github.com/YuriyLisovskiy/codegen/generators"
	"github.com/YuriyLisovskiy/codegen/parser"
	"os"
	"path/filepath"
)

var (
	langPtr   = flag.String("l", defaultLang, "language")
	xmlPtr    = flag.String("f", "", "file")
	xmlUrlPtr = flag.String("u", "", "file url")
	spacesPtr = flag.Bool("s", false, "use spaces")
	outputDir = "out"
)

func parseFileByFormat(body []byte, fileName string) (generators.Package, error) {
	var pkg = generators.Package{}

	fileFormat, err := parser.GetFileFormat(fileName)
	if err != nil {
		return pkg, err
	}

	switch fileFormat {
	case "xml":
		xml.Unmarshal(body, &pkg)
	case "json":
		json.Unmarshal(body, &pkg)
	case "yml":
		yaml.Unmarshal(body, &pkg)
	default:
		return pkg, errors.New(fmt.Sprintf("Invalid format of '%s' file.", fileName))
	}
	return pkg, nil
}

func getSerializedData(url string, fileName string) ([]byte, error) {
	if url != "" {
		fileName = url
		return parser.Download(url)
	} else {
		return parser.Read(fileName)
	}
}

func putDataToFiles(fileContextMap map[string]string, lang string) error {
	ext := parser.GetExtension(lang)

	os.MkdirAll(outputDir, os.ModePerm)
	for key, val := range fileContextMap {
		err := parser.Write(filepath.Join(outputDir, key+ext), val)
		if err != nil {
			return err
		}
	}
	fmt.Println("Generated successfully.")

	return nil
}

func putDataToStdout(fileContextMap map[string]string, lang string) error {
	ext := parser.GetExtension(lang)
	comment := generators.Generators[lang].Comment
	filename := ""
	for key, val := range fileContextMap {
		if comment != "" {
			filename = fmt.Sprintf(comment, filepath.Join(outputDir, key+ext))
		}
		//fmt.Printf("%s\n%s\n", comment, val)
		fmt.Print(filename)
		fmt.Printf("\n%s\n", val)
	}
	return nil
}

func execute() error {
	flag.Parse()
	language, fileName, url, useSpaces := *langPtr, *xmlPtr, *xmlUrlPtr, *spacesPtr

	err := parser.ValidateArgs(language, fileName, url)
	if err != nil {
		return err
	}

	byteContext, err := getSerializedData(url, fileName)
	if err != nil {
		return err
	}

	generator, err := generators.GetGenerator(language)
	if err != nil {
		return nil
	}

	pkg := generators.Package{UseSpaces: useSpaces}

	pkg, err = parseFileByFormat(byteContext, fileName)
	//if err != nil {
	//	return nil
	//}

	pkg = defaultPkg
	if len(pkg.Classes) == 0 && len(pkg.Classes) == 0 && len(pkg.Classes) == 0 {
		return errors.New("There is nothing to generate, file '" + fileName + "' is empty or incorrect.")
	}

	fileContextMap := generator.Generate(pkg)

	putDataToStdout(fileContextMap, language)

	return nil
}
