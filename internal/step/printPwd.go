package step

import (
	"os"
	"fmt"
)

func init () {
	step := &PrintPwd{}
	Register(step)
}

type PrintPwd struct {}

func (*PrintPwd) Validate(ctx Context) error {
	return nil
}

func (*PrintPwd) Execute(ctx Context) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(pwd)

	return nil
}

func (*PrintPwd) Code() string {
	return "printPwd"
}

func (*PrintPwd) Name() string {
	return "Print PWD"
}

func (*PrintPwd) Description() string {
	return "Print to stdout current working directory"
}

func (*PrintPwd) Params() map[string]string {
	return map[string]string{}
}