package parser

type Field struct {
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
	Name       string `xml:"name"`
	Return     string `xml:"return"`
	Const      bool   `xml:"const"`
	Parameters []Parameter
}

type Constructor struct {
	Explicit   bool `xml:"explicit"`
	Const      bool `xml:"const"`
	Parameters []Parameter
}

type Class struct {
	Fields       []Field
	Methods      []Function
	Constructors []Constructor
	Classes      []Class
}

