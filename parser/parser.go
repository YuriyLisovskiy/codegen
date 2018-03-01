package parser

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"os"
	"io"
	"strings"
	"encoding/json"
	"gopkg.in/yaml.v2"
)


type Field struct {
	Name    string `xml:"name" json:"name" yml:"name"`
	Type    string `xml:"type" json:"type" yml:"type"`
	Const   bool   `xml:"const" json:"const" yml:"const"`
	Pointer bool   `xml:"pointer" json:"pointer" yml:"pointer"`
	Default string `xml:"default" json:"default" yml:"default"`
	Access  string `xml:"access" json:"access" yml:"access"`
	Static  bool   `xml:"static" json:"static" yml:"static"`
}

type Parameter struct {
	Name    string `xml:"name" json:"name" yml:"name"`
	Type    string `xml:"type" json:"type" yml:"type"`
	Pass    string `xml:"pass" json:"pass" yml:"pass"`
	Const   bool   `xml:"const" json:"const" yml:"const"`
	Default string `xml:"default" json:"default" yml:"default"`
}

type Method struct {
	Name       string      `xml:"name" json:"name" yml:"name"`
	Return     string      `xml:"return" json:"return" yml:"return"`
	Access     string      `xml:"access" json:"access" yml:"access"`
	Const      bool        `xml:"const" json:"const" yml:"const"`
	Static     bool        `xml:"static" json:"static" yml:"static"`
	Parameters []Parameter `xml:"parameters>parameter" json:"parameters" yml:"parameters"`
}

type Constructor struct {
	Explicit   bool        `xml:"explicit" json:"explicit" yml:"explicit"`
	Const      bool        `xml:"const" json:"const" yml:"const"`
	Parameters []Parameter `xml:"parameters>parameter" json:"parameters" yml:"parameters"`
}

type Parent struct {
	Name   string `xml:"name" json:"name" yml:"name"`
	Access string `xml:"access" json:"access" yml:"access"`
}

type Class struct {
	Name         string        `xml:"name" json:"name" yml:"name"`
	Fields       []Field       `xml:"fields>field" json:"fields" yml:"fields"`
	Methods      []Method      `xml:"methods>method" json:"methods" yml:"methods"`
	Constructors []Constructor `xml:"constructors>constructor" json:"constructors" yml:"constructors"`
	Classes      []Class       `xml:"classes>class" json:"classes" yml:"classes"`
	Parent       Parent        `xml:"parent" json:"parent" yml:"parent"`
	Access       string        `xml:"access" json:"access" yml:"access"`
}

type Package struct {
	Classes   []Class  `xml:"class" json:"classes" yml:"classes"`
	Functions []Method `xml:"function" json:"functions" yml:"functions"`
	Variables []Field  `xml:"variable" json:"variables" yml:"variables"`
	UseSpaces bool
}

func Read(name string) ([]byte, error) {
	xmlFile, err := ioutil.ReadFile(name)
	if err != nil {
		return []byte(""), err
	}
	return xmlFile, nil
}

func Write(path, fileContext string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, strings.NewReader(fileContext))
	if err != nil {
		return err
	}
	return nil
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

func ParseXml(file []byte) Package {
	var obj Package
	xml.Unmarshal(file, &obj)
	return obj
}

func ParseJson(file []byte) Package {
	var obj Package
	json.Unmarshal(file, &obj)
	return obj
}

func ParseYaml(file []byte) Package {
	var obj Package
	yaml.Unmarshal(file, &obj)
	return obj
}
