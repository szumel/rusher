package macro

import (
	"fmt"
	"io"
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

type filesystemMock struct{}

func (*filesystemMock) open(name string) (file, error) {
	return &fileMock{}, nil
}

type fileMock struct {
	done bool
}

func (f *fileMock) Read(p []byte) (n int, err error) {
	if f.done {
		return 0, io.EOF
	}
	for i, b := range []byte(expectedSchema) {
		p[i] = b
	}

	f.done = true

	return len(p), nil
}

func (*fileMock) Close() error {
	return nil
}
