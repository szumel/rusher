package step

import (
	"errors"
	"os"
	"fmt"
)

func init() {
	Register(&RemoveDir{})
}

type RemoveDir struct {}

func (*RemoveDir) Execute(ctx Context) error {
	fmt.Println("Removing " + ctx.Params()["dir"])
	return os.RemoveAll(ctx.Params()["dir"])
}

func (*RemoveDir) Code() string {
	return "removeDir"
}

func (*RemoveDir) Name() string {
	return "Remove Dir"
}

func (*RemoveDir) Description() string {
	return "Removes given direcory"
}

func (*RemoveDir) Params() map[string]string {
	params := map[string]string {
		"dir": "direcory path which will be removed",
	}

	return params
}

func (*RemoveDir) Validate(ctx Context) error {
	if ctx.Params()["dir"] == "" {
		return errors.New("RemoveDir: don't your forget to provide dir param?")
	}

	return nil
}
