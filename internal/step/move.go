package step

import (
	"os"
	"fmt"
)

func init() {
	Register(&Move{})
}

type Move struct{}

func (*Move) Execute(ctx Context) error {
	source := ctx.Params()["source"]
	dst := ctx.Params()["dst"]

	fmt.Printf("Moving %s to %s \n", source, dst)

	return os.Rename(source, dst)
}

func (*Move) Code() string {
	return "move"
}

func (*Move) Name() string {
	return "Move"
}

func (*Move) Description() string {
	return "Move file/dir to new location with new name"
}

func (*Move) Params() map[string]string {
	params := map[string]string{
		"source": "source file/dir path to move",
		"dst":    "destination file/dir path",
	}

	return params
}

func (m *Move) Validate(ctx Context) error {
	source := ctx.Params()["source"]
	dst := ctx.Params()["dst"]

	if source == "" || dst == "" {
		return NewError(m.Name(), "source or dst must be provided")
	}

	if _, err := os.Stat(source); os.IsNotExist(err) {
		return NewError(m.Name(), "given source path does not exist")
	}

	return nil
}
