package parser

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"os"
	"io"
	"strings"
	"encoding/json"
)


type Field struct {
	Name    string `xml:"name" json:"name"`
	Type    string `xml:"type" json:"type"`
	Const   bool   `xml:"const" json:"const"`
	Pointer bool   `xml:"pointer" json:"pointer"`
	Default string `xml:"default" json:"default"`
	Access  string `xml:"access" json:"access"`
	Static  bool   `xml:"static" json:"static"`
}

type Parameter struct {
	Name    string `xml:"name" json:"name"`
	Type    string `xml:"type" json:"type"`
	Pass    string `xml:"pass" json:"pass"`
	Const   bool   `xml:"const" json:"const"`
	Default string `xml:"default" json:"default"`
}

type Method struct {
	Name       string      `xml:"name" json:"name"`
	Return     string      `xml:"return" json:"return"`
	Access     string      `xml:"access" json:"access"`
	Const      bool        `xml:"const" json:"const"`
	Static     bool        `xml:"static" json:"static"`
	Parameters []Parameter `xml:"parameters>parameter" json:"parameters"`
}

type Constructor struct {
	Explicit   bool        `xml:"explicit" json:"explicit"`
	Const      bool        `xml:"const" json:"const"`
	Parameters []Parameter `xml:"parameters>parameter" json:"parameters"`
}

type Parent struct {
	Name   string `xml:"name" json:"name"`
	Access string `xml:"access" json:"access"`
}

type Class struct {
	Name         string        `xml:"name" json:"name"`
	Fields       []Field       `xml:"fields>field" json:"fields"`
	Methods      []Method      `xml:"methods>method" json:"methods"`
	Constructors []Constructor `xml:"constructors>constructor" json:"constructors"`
	Classes      []Class       `xml:"classes>class" json:"classes"`
	Parent       Parent        `xml:"parent" json:"parent"`
	Access       string        `xml:"access" json:"access"`
}

type Package struct {
	Classes   []Class  `xml:"class" json:"classes"`
	Functions []Method `xml:"function" json:"functions"`
	Variables []Field  `xml:"variable" json:"variables"`
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
