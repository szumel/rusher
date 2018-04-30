package macro

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

var (
	expectedSchema Schema = `<?xml version="1.0"?>
<macro v="1.0.0">
    <step code="printPwd"/>
    <step code="changeCwd" dir="/xx"/>
</macro>`

	stringExpectedSchema = removeWhites(string(expectedSchema))
)

//@todo test GIT
func TestLoad(t *testing.T) {
	sources := []string{
		"macro.xml",
	}

	for _, source := range sources {
		loader := Loader{filesystem: &filesystemMock{}, git: &osFilesystem{}}
		schema, err := loader.Load(source)
		if err != nil {
			log.Fatal(err)
		}

		stringSchema := removeWhites(string(schema))
		fmt.Println("exx", stringExpectedSchema, len(stringExpectedSchema))
		fmt.Println("got", stringSchema, len(schema))
		if stringSchema != stringExpectedSchema {
			log.Fatal("macro.Load failed: loaded schema is not same as expected")
		}
	}
}

func removeWhites(s string) string {
	s = strings.Replace(string(s), " ", "", -1)
	s = strings.Replace(s, "\n", "", -1)

	return s
}

type filesystemMock struct{}

func (*filesystemMock) open(name string) (file, error) {
	return &fileMock{}, nil
}

type fileMock struct{}

func (*fileMock) Read(p []byte) (n int, err error) {
	r := strings.NewReader(stringExpectedSchema)

	return r.Read(p)
}

func (*fileMock) Close() error {
	return nil
}
