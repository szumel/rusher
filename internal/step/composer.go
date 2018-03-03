package step

import (
	"os/exec"
	"os"
	"fmt"
)

func init() {
	Register(&ComposerInstall{})
}

type ComposerInstall struct{}

func (*ComposerInstall) Execute(ctx Context) error {
	fmt.Println("Executing composer install")
	cmd := exec.Command("composer", "install")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (*ComposerInstall) Code() string {
	return "composerInstall"
}

func (*ComposerInstall) Name() string {
	return "Composer install"
}

func (*ComposerInstall) Description() string {
	return "Install composer dependencies"
}

func (*ComposerInstall) Params() map[string]string {
	return map[string]string{}
}

func (*ComposerInstall) Validate(ctx Context) error {
	return nil
}
