package step

import (
	"os/exec"
	"os"
)

type aptInstall struct {}

func (*aptInstall) Execute(ctx Context) error {
	cmd := exec.Command("apt", "install", ctx.Params()["package"], "-"+ctx.Params()["accept"])
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

func (*aptInstall) Code() string {
	return "aptInstall"
}

func (*aptInstall) Name() string {
	return "aptInstall"
}

func (*aptInstall) Description() string {
	return "install package via apt get"
}

func (*aptInstall) Params() map[string]string {
	return map[string]string{
		"package": "package to install",
		"accept": "accept all dependencies to install [y/n]",
	}
}

func (a *aptInstall) Validate(ctx Context) error {
	if ctx.Params()["package"] == "" {
		return NewError(a.Name(), "package is required")
	}

	if ctx.Params()["accept"] != "y" && ctx.Params()["accept"] != "n" {
		return NewError(a.Name(), "accept accepts y or n value")
	}

	return nil
}
