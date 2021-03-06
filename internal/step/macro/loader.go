package macro

import (
	"io"
	"io/ioutil"
	"os"
)

type Schema string

type file interface {
	io.Closer
	io.Reader
}

type storage interface {
	open(name string) (file, error)
}

func CreateLoader() *LoaderImpl {
	return &LoaderImpl{
		filesystem: &osFilesystem{},
		git:        &osFilesystem{},
	}
}

type Loader interface {
	Load(source string) (Schema, error)
}

type LoaderImpl struct {
	git        storage
	filesystem storage
}

func (l *LoaderImpl) Load(source string) (Schema, error) {
	stype := resolveSourceType(source)
	var schema Schema
	switch stype {
	case sourceFile:
		s, err := l.loadFile(source)
		schema = s
		if err != nil {
			return "", err
		}
		break
	}

	return schema, nil
}

func (l *LoaderImpl) loadFile(source string) (Schema, error) {
	f, err := l.filesystem.open(source)
	if err != nil {
		return "", err
	}

	defer f.Close()

	s, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	return Schema(s), nil
}

type osFilesystem struct{}

func (*osFilesystem) open(name string) (file, error) {
	return os.Open(name)
}
