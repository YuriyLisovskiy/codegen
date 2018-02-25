package parser

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type Field struct {
	Name    string `xml:"name"`
	Type    string `xml:"type"`
	Const   bool   `xml:"const"`
	Pointer bool   `xml:"pointer"`
	Default string `xml:"default"`
	Access  string `xml:"access"`
	Static  bool   `xml:"static"`
}

type Parameter struct {
	Name    string `xml:"name"`
	Type    string `xml:"type"`
	Pass    string `xml:"pass"`
	Const   bool   `xml:"const"`
	Default string `xml:"default"`
}

type Method struct {
	Name       string      `xml:"name"`
	Return     string      `xml:"return"`
	Access     string      `xml:"access"`
	Const      bool        `xml:"const"`
	Static     bool        `xml:"static"`
	Parameters []Parameter `xml:"parameters>parameter"`
}

type Constructor struct {
	Explicit   bool        `xml:"explicit"`
	Const      bool        `xml:"const"`
	Parameters []Parameter `xml:"parameters>parameter"`
}

type Parent struct {
	Name string `xml:"name"`
	Access string `xml:"access"`
}

type Class struct {
	Name         string        `xml:"name"`
	Fields       []Field       `xml:"fields>field"`
	Methods      []Method      `xml:"methods>method"`
	Constructors []Constructor `xml:"constructors>constructor"`
	Classes      []Class       `xml:"classes>class"`
	Parent       Parent        `xml:"parent"`
	Access       string        `xml:"access"`
}

type Xml struct {
	Classes   []Class  `xml:"classes>class"`
	Functions []Method `xml:"functions>function"`
	Variables []Field  `xml:"variables>variable"`
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
