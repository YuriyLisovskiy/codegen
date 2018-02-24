package parser

import (
	"encoding/xml"
	"io/ioutil"
	"fmt"
)

type Variable struct {
	Name    string `xml:"name"`
	Type    string `xml:"type"`
	Const   bool   `xml:"const"`
	Pointer bool   `xml:"pointer"`
	Default string `xml:"default"`
}

type Parameter struct {
	Name    string `xml:"name"`
	Type    string `xml:"type"`
	Pass    string `xml:"pass"`
	Const   bool   `xml:"const"`
	Default string `xml:"default"`
}

type Function struct {
	Name       string      `xml:"name"`
	Return     string      `xml:"return"`
	Const      bool        `xml:"const"`
	Parameters []Parameter `xml:"parameters>parameter"`
}

type Constructor struct {
	Explicit   bool        `xml:"explicit"`
	Const      bool        `xml:"const"`
	Parameters []Parameter `xml:"parameters>parameter"`
}

type Class struct {
	Name         string        `xml:"name"`
	Fields       []Variable    `xml:"fields>field"`
	Methods      []Function    `xml:"methods>method"`
	Constructors []Constructor `xml:"constructors>constructor"`
	Classes      []Class       `xml:"classes>class"`
	Parents      []string      `xml:"parents>parent"`
}

type Xml struct {
	Classes []Class `xml:"classes>class"`
	Functions []Function `xml:"functions>function"`
	Variables []Variable `xml:"variables>variable"`
}

func read(name string) []byte {
	xmlFile, errors := ioutil.ReadFile(name)
	if errors != nil {
		fmt.Println(errors)
	}
	return xmlFile
}

func parse(file []byte) Xml {
	var obj Xml
	xml.Unmarshal(file, &obj)
	return obj
}
