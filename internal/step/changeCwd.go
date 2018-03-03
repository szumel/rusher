package step

import (
	"errors"
	"fmt"
	"os"
)

func init() {
	step := &ChangeCwd{}
	Register(step)
}

type ChangeCwd struct{}

func (*ChangeCwd) Execute(ctx Context) error {
	var newCwd string
	if ctx.Params()["dir"] == "{projectPath}" {
		newCwd = ctx.ProjectPath()
	} else {
		newCwd = ctx.Params()["dir"]
	}
	fmt.Println("changing current working directory to " + newCwd)

	return os.Chdir(newCwd)
}

func (*ChangeCwd) Code() string {
	return "changeCwd"
}

func (*ChangeCwd) Name() string {
	return "Change Current Working Directory"
}

func (*ChangeCwd) Description() string {
	return "Changing current working directory to new"
}

func (*ChangeCwd) Params() map[string]string {
	params := map[string]string{
		"dir": "New current working directory. There are predefined values you can use: {projectPath}",
	}

	return params
}

func (*ChangeCwd) Validate(ctx Context) error {
	if ctx.Params()["dir"] == "" {
		return errors.New("ChangeCwd: dir param must be provided")
	}

	return nil
}
