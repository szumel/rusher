package macro

import (
	"encoding/xml"
	"log"
	"testing"
)

func TestResolveSourceType(t *testing.T) {
	sources := map[string]string{
		"macro.xml": sourceFile,
		"xxx":       sourceFile,
		"http://github.com/macro.git": sourceGitHttp,
		"git@github.com/macro.git":    sourceGitSsh,
	}

	for source, expectedType := range sources {
		st := resolveSourceType(source)
		if st != expectedType {
			t.Fatalf("macro.resolveSourceType failed: expected %s type for %s source, got %s", expectedType, source, st)
		}
	}
}

func TestCreate(t *testing.T) {
	var in Schema = `<macro v="1.0.0">
				<step code="printPwd"/>
				<step code="changeCwd" dir="/xx"/>
			</macro>`

	m, err := Create(in)
	if err != nil {
		log.Fatal(err)
	}

	if m.Version != "1.0.0" {
		log.Fatalf("macro.Create failed: version expected as %s got %s", "1.0.0", m.Version)
	}

	if len(m.Steps) != 2 {
		log.Fatalf("macro.Create failed: expected steps length %d got %d", 2, len(m.Steps))
	}

	expectedSteps := []Step{
		{Code: "printPwd"},
		{Code: "changeCwd", Params: []xml.Attr{xml.Attr{Name: xml.Name{Local: "dir"}, Value: "/xx"}}},
	}

	for index, step := range m.Steps {
		if expectedSteps[index].Code != step.Code {
			log.Fatalf("macro.Create failed: steps sort is invalid, expected %s got %s", expectedSteps[index].Code, step.Code)
		}
	}
}
