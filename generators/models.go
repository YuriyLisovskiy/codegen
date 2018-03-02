package generators

type (
	Field struct {
		Name    string `xml:"name" json:"name" yml:"name"`
		Type    string `xml:"type" json:"type" yml:"type"`
		Const   bool   `xml:"const" json:"const" yml:"const"`
		Pointer bool   `xml:"pointer" json:"pointer" yml:"pointer"`
		Default string `xml:"default" json:"default" yml:"default"`
		Access  string `xml:"access" json:"access" yml:"access"`
		Static  bool   `xml:"static" json:"static" yml:"static"`
	}

	Parameter struct {
		Name    string `xml:"name" json:"name" yml:"name"`
		Type    string `xml:"type" json:"type" yml:"type"`
		Pass    string `xml:"pass" json:"pass" yml:"pass"`
		Const   bool   `xml:"const" json:"const" yml:"const"`
		Default string `xml:"default" json:"default" yml:"default"`
	}

	Method struct {
		Name       string      `xml:"name" json:"name" yml:"name"`
		Return     string      `xml:"return" json:"return" yml:"return"`
		Access     string      `xml:"access" json:"access" yml:"access"`
		Const      bool        `xml:"const" json:"const" yml:"const"`
		Static     bool        `xml:"static" json:"static" yml:"static"`
		Parameters []Parameter `xml:"parameters>parameter" json:"parameters" yml:"parameters"`
	}

	Constructor struct {
		Explicit   bool        `xml:"explicit" json:"explicit" yml:"explicit"`
		Const      bool        `xml:"const" json:"const" yml:"const"`
		Parameters []Parameter `xml:"parameters>parameter" json:"parameters" yml:"parameters"`
	}

	Parent struct {
		Name   string `xml:"name" json:"name" yml:"name"`
		Access string `xml:"access" json:"access" yml:"access"`
	}

	Class struct {
		Name         string        `xml:"name" json:"name" yml:"name"`
		Fields       []Field       `xml:"fields>field" json:"fields" yml:"fields"`
		Methods      []Method      `xml:"methods>method" json:"methods" yml:"methods"`
		Constructors []Constructor `xml:"constructors>constructor" json:"constructors" yml:"constructors"`
		Classes      []Class       `xml:"classes>class" json:"classes" yml:"classes"`
		Parent       Parent        `xml:"parent" json:"parent" yml:"parent"`
		Access       string        `xml:"access" json:"access" yml:"access"`
	}

	Package struct {
		Name      string   `xml:"name" json:"name" yml:"name"`
		Classes   []Class  `xml:"classes>class" json:"classes" yml:"classes"`
		Functions []Method `xml:"functions>function" json:"functions" yml:"functions"`
		Variables []Field  `xml:"variables>variable" json:"variables" yml:"variables"`
		UseSpaces bool     `xml:"use_spaces" json:"use_spaces" yml:"use_spaces"`
	}
)
