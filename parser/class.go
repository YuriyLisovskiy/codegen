package parser


type Class struct {
	Fields       []struct {
		Name string `xml:"name"`
		Type string `xml:"type"`
		Const bool `xml:"const"`
		Pointer bool `xml:"pointer"`
		Default string `xml:"default"`
	}
	Methods      []struct {
		Name string `xml:"name"`
		Return string `xml:"return"`
		Const bool `xml:"const"`
		Parameters [] struct {
			Name string `xml:"name"`
			Type string `xml:"type"`
			Pass string `xml:"pass"`
			Const bool `xml:"const"`
			Default string `xml:"default"`
		}
	}
	Constructors []struct {
		Explicit bool `xml:"explicit"`
		Const bool `xml:"const"`
		Parameters [] struct {
			Name string `xml:"name"`
			Type string `xml:"type"`
			Pass string `xml:"pass"`
			Const bool `xml:"const"`
			Default string `xml:"default"`
		}
	}
	Classes      []Class
}

