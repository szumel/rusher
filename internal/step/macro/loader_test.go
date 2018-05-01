package macro

import (
	"log"
	"strings"
	"testing"
)

//@todo test GIT
func TestLoad(t *testing.T) {
	sources := []string{
		"macro.xml",
	}

	for _, source := range sources {
		loader := LoaderImpl{filesystem: &filesystemMock{file: &fileMock{data: []byte(expectedSchema)}}, git: &osFilesystem{}}
		schema, err := loader.Load(source)
		if err != nil {
			log.Fatal(err)
		}

		stringSchema := removeWhites(string(schema))
		stringExpectedSchema := removeWhites(string(expectedSchema))

		if stringSchema != stringExpectedSchema {
			log.Fatal("macro.Load failed: loaded schema is not same as expected")
		}
	}
}

func removeWhites(s string) string {
	s = strings.Replace(string(s), " ", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\x00", "", -1)

	return s
}
