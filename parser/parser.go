package parser

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
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

func Read(name string) ([]byte, error) {
	xmlFile, err := ioutil.ReadFile(name)
	if err != nil {
		return []byte(""), err
	}
	return xmlFile, nil
}

func Download(url string) ([]byte, error) {
	response, err := http.Get(url)
	result := []byte("")
	if err != nil {
		return result, err
	}
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return result, err
		}
		result = bodyBytes
	}
	return result, nil
}

func Parse(file []byte) Xml {
	var obj Xml
	xml.Unmarshal(file, &obj)
	return obj
}
