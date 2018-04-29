package step

import (
	"errors"
	"fmt"
)

type PrintString struct{}

func (*PrintString) Execute(ctx Context) error {
	for _, value := range ctx.Params() {
		fmt.Println(value)
	}

	return nil
}

func (*PrintString) Code() string {
	return "printString"
}

func (*PrintString) Name() string {
	return "Print String"
}

func (*PrintString) Description() string {
	return "Printing to stdout provided string. \n String should be provided as param 'text'."
}

func (*PrintString) Params() map[string]string {
	return map[string]string{"text": "string"}
}

func (*PrintString) Validate(ctx Context) error {
	_, ok := ctx.Params()["text"]
	if !ok {
		return errors.New("PrintString: text parameter must be provided")
	}

	return nil
}
