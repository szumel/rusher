package step

import (
	"fmt"
	"os"
	"rusher/internal/platform/risk"
)

func init() {
	Register(&RemoveDir{})
}

type RemoveDir struct{}

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
	return "Removes given directory"
}

func (*RemoveDir) Params() map[string]string {
	params := map[string]string{
		"dir": "directory path which will be removed",
	}

	return params
}

func (r *RemoveDir) Validate(ctx Context) error {
	if ctx.Params()["dir"] == "" {
		return NewError(r.Name(), "don't your forget to provide dir param?")
	}

	err := risk.DirCheck(ctx.Params()["dir"])
	if err != nil {
		return NewError(r.Name(), ctx.Params()["dir"]+" directory has been detected as risky")
	}

	return nil
}
