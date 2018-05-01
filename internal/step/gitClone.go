package step

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4"
	gitssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"io/ioutil"
)

type GitClone struct{}

func (*GitClone) Execute(ctx Context) error {
	key, err := ioutil.ReadFile(ctx.Params()["key"])
	if err != nil {
		return err
	}

	signer, err := ssh.ParsePrivateKey([]byte(key))
	if err != nil {
		return err
	}

	auth := &gitssh.PublicKeys{User: "git", Signer: signer}

	fmt.Println("cloning " + ctx.Params()["origin"] + " to " + ctx.Params()["dir"])
	_, err = git.PlainClone(ctx.Params()["dir"], false, &git.CloneOptions{
		URL:               ctx.Params()["origin"],
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Auth:              auth,
	})

	fmt.Println("cloned " + ctx.Params()["origin"])

	return err
}

func (*GitClone) Code() string {
	return "gitClone"
}

func (*GitClone) Name() string {
	return "Git clone"
}

func (*GitClone) Description() string {
	return "Cloning git repository"
}

func (*GitClone) Params() map[string]string {
	return map[string]string{
		"origin": "repository origin path",
		"dir":    "to which location clone",
		"key":    "path to ssh key",
	}
}

func (*GitClone) Validate(ctx Context) error {
	//@todo check if key exists?
	if ctx.Params()["origin"] == "" {
		return errors.New("GitClone: origin param must be provided")
	}

	if ctx.Params()["dir"] == "" {
		return errors.New("GitClone: dir param must be provided")
	}

	if ctx.Params()["key"] == "" {
		return errors.New("GitClone: key param must be provided")
	}

	return nil
}
