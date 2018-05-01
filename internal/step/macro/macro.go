package macro

import (
	"encoding/xml"
	"net/url"
	"strings"
)

const (
	sourceFile    = "file"
	sourceGitHttp = "gitHttp"
	sourceGitSsh  = "gitSsh"
)

//Step represents each step nested in Macro
type Step struct {
	Code   string     `xml:"code,attr"`
	Params []xml.Attr `xml:",any,attr"`
}

//Macro represents steps aggregate
//Each macro is defined exactly in one file
type Macro struct {
	Version string `xml:"v,attr"`
	Steps   []Step `xml:"step"`
}

func Create(source Schema) (*Macro, error) {
	var macro Macro
	err := xml.Unmarshal([]byte(source), &macro)
	if err != nil {
		return nil, err
	}

	return &macro, nil
}

func resolveSourceType(source string) string {
	stype := sourceFile
	if strings.Contains(source, ".git") {
		if _, err := url.ParseRequestURI(source); err == nil {
			stype = sourceGitHttp
		} else {
			stype = sourceGitSsh
		}
	}

	return stype
}
