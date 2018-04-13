package step

import (
	"github.com/szumel/rusher/internal/step"
	"os"
	"testing"
)

const dir = "/"

func TestChangeCwd(t *testing.T) {
	s := step.ChangeCwd{}
	ctx := ctxMock{}

	err := s.Execute(&ctx)
	if err != nil {
		t.Fatal(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	if cwd != dir {
		t.Fatalf("Expected current working directory %s, have %s", dir, cwd)
	}
}

type ctxMock struct{}

func (*ctxMock) ProjectPath() string {
	return ""
}

func (*ctxMock) Params() map[string]string {
	return map[string]string{
		"dir": dir,
	}
}

func (*ctxMock) Globals() map[string]string {
	return nil
}
