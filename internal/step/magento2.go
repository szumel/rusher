package step

import (
	"os"
	"os/exec"
)

type Magento2EnableModules struct{}

func (*Magento2EnableModules) Execute(ctx Context) error {
	cmd := exec.Command("bin/magento", "module:enable", "--all")
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func (*Magento2EnableModules) Code() string {
	return "magento2EnableModules"
}

func (*Magento2EnableModules) Name() string {
	return "Magento2 enable modules"
}

func (*Magento2EnableModules) Description() string {
	return "Enable all modules in magento 2 instance"
}

func (*Magento2EnableModules) Params() map[string]string {
	return map[string]string{}
}

func (*Magento2EnableModules) Validate(ctx Context) error {
	return nil
}

type Magento2SetupUpgrade struct{}

func (*Magento2SetupUpgrade) Execute(ctx Context) error {
	cmd := exec.Command("bin/magento", "setup:upgrade")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (*Magento2SetupUpgrade) Code() string {
	return "magento2SetupUpgrade"
}

func (*Magento2SetupUpgrade) Name() string {
	return "Magento 2 setup upgrade"
}

func (*Magento2SetupUpgrade) Description() string {
	return "Executes setup:upgrade command on magento 2 intance"
}

func (*Magento2SetupUpgrade) Params() map[string]string {
	return map[string]string{}
}

func (*Magento2SetupUpgrade) Validate(ctx Context) error {
	return nil
}

type Magento2Compile struct{}

func (*Magento2Compile) Execute(ctx Context) error {
	cmd := exec.Command("bin/magento", "setup:di:compile")
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func (*Magento2Compile) Code() string {
	return "magento2Compile"
}

func (*Magento2Compile) Name() string {
	return "Magento 2 compile"
}

func (*Magento2Compile) Description() string {
	return "Compiles magento 2 instance"
}

func (*Magento2Compile) Params() map[string]string {
	return map[string]string{}
}

func (*Magento2Compile) Validate(ctx Context) error {
	return nil
}
