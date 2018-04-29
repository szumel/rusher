package macro

import (
	"io/ioutil"
)

type Schema string

func Load(source string) (Schema, error) {
	stype := resolveSourceType(source)
	var schema Schema
	switch stype {
	case sourceFile:
		s, err := loadFile(source)
		schema = s
		if err != nil {
			return "", err
		}
		break
	}

	return schema, nil
}

func loadFile(source string) (Schema, error) {
	f, err := ioutil.ReadFile(source)
	if err != nil {
		return "", err
	}

	return Schema(f), nil
}
