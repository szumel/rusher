package step

import "fmt"

type Error struct {
	stepName string
	msg      string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.stepName, e.msg)
}

func NewError(stepName string, msg string) error {
	return &Error{stepName: stepName, msg: msg}
}
