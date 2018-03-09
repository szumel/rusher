package step

import (
	"errors"
	"fmt"
	"os"
	"rusher/internal/platform/rollback"
)

func init() {
	step := &Symlink{}
	Register(step)
	rollback.Subscribe(step)
}

type Symlink struct {
	ctx Context
}

func (s *Symlink) Rollback() error {
	if _, err := os.Stat(s.ctx.Params()["target"]); !os.IsNotExist(err) {
		fmt.Println("removing symlink " + s.ctx.Params()["source"])
		return os.Remove(s.ctx.Params()["source"])
	}

	return nil
}

func (s *Symlink) Execute(ctx Context) error {
	fmt.Println(fmt.Sprintf("creating symlink %s to %s", ctx.Params()["source"], ctx.Params()["target"]))

	return os.Symlink(ctx.Params()["target"], ctx.Params()["source"])
}

func (*Symlink) Code() string {
	return "symlink"
}

func (*Symlink) Name() string {
	return "Symlink"
}

func (*Symlink) Description() string {
	return "create symlink in context of current working direcory"
}

func (*Symlink) Params() map[string]string {
	params := map[string]string{}
	params["source"] = "link source"
	params["target"] = "link target"

	return params
}

func (s *Symlink) Validate(ctx Context) error {
	s.ctx = ctx
	if ctx.Params()["source"] == "" || ctx.Params()["target"] == "" {
		return errors.New("Symlink: source and target must be provided")
	}

	return nil
}
