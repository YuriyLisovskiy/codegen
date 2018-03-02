package generators

import (
	"encoding/json"
	"encoding/xml"
	"gopkg.in/yaml.v2"
)

type (
	JsonGenerator struct{}
	XmlGenerator  struct{}
	YamlGenerator struct{}
)

func (gen XmlGenerator) Generate(pkg Package) map[string]string {
	bytes, _ := xml.MarshalIndent(pkg, "", getIndent(!pkg.UseSpaces, 2))
	return map[string]string{
		pkg.Name: "<?xml version=\"1.0\" encoding=\"UTF-8\" ?>\n" + string(bytes),
	}
}
func (gen JsonGenerator) Generate(pkg Package) map[string]string {
	bytes, _ := json.MarshalIndent(pkg, "", getIndent(!pkg.UseSpaces, 2))
	return map[string]string{
		pkg.Name: string(bytes),
	}
}
func (gen YamlGenerator) Generate(pkg Package) map[string]string {
	bytes, _ := yaml.Marshal(pkg)
	return map[string]string{
		pkg.Name: string(bytes),
	}
}

func (gen XmlGenerator) generateField(field Field) string    { return "" }
func (gen XmlGenerator) generateMethod(method Method) string { return "" }
func (gen XmlGenerator) generateClass(class Class) string    { return "" }

func (gen JsonGenerator) generateField(field Field) string    { return "" }
func (gen JsonGenerator) generateMethod(method Method) string { return "" }
func (gen JsonGenerator) generateClass(class Class) string    { return "" }

func (gen YamlGenerator) generateField(field Field) string    { return "" }
func (gen YamlGenerator) generateMethod(method Method) string { return "" }
func (gen YamlGenerator) generateClass(class Class) string    { return "" }
