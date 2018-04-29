package macro

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

//@todo GIT
func TestLoad(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	sources := []string{
		wd + "/macro.xml",
	}
	var expectedSchema Schema = `<?xml version="1.0"?>

<macro v="1.0.0">
    <step code="printPwd"/>
    <step code="changeCwd" dir="/xx"/>
</macro>`

	for _, source := range sources {
		fmt.Println(source)
		schema, err := Load(source)
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

	return s
}
