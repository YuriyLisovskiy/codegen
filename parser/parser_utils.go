package parser

import (
	"errors"
	"github.com/YuriyLisovskiy/codegen/generators"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

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

func ValidateArgs(lang, file, url string) error {
	if lang == "" {
		return errors.New("specify language (-l) flag")
	}
	if url == "" && file == "" {
		return errors.New("specify file path (-f) or url path (-u) flag")
	}
	if file != "" && url != "" {
		return errors.New("do not use both -f and -u flags at the same time")
	}
	return nil
}

func GetExtension(language string) string {
	language = generators.NormalizeLang(language)
	ext := generators.Generators[language].Extension
	if ext != "" {
		ext = "." + ext
	}
	return ext
}

func GetFileFormat(name string) (string, error) {
	arr := strings.Split(name, ".")
	if len(arr) > 0 {
		return arr[len(arr)-1], nil
	}
	return "", errors.New("invalid input file")
}
