package step

import (
	"fmt"
	"io"
	"os"
)

func init() {
	Register(&copyFiles{})
}

type copyFiles struct{}

func (*copyFiles) Validate(ctx Context) error {
	fromPath := ctx.Params()["from"]
	_, err := os.Stat(fromPath)
	if err != nil {
		return err
	}

	return nil
}

func (*copyFiles) Execute(ctx Context) error {
	fromPath := ctx.Params()["from"]
	toPath := ctx.Params()["to"]

	fmt.Println("Copying from " + fromPath + " to " + toPath)
	src, _ := os.Stat(fromPath)
	switch mode := src.Mode(); {
	case mode.IsDir():
		err := copyDir(fromPath, toPath)
		if err != nil {
			return err
		}
	case mode.IsRegular():
		err := copyFile(fromPath, toPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (*copyFiles) Code() string {
	return "copyFiles"
}

func (*copyFiles) Name() string {
	return "Copy Files"
}

func (*copyFiles) Description() string {
	return "Copy files and folders from place to place"
}

func (*copyFiles) Params() map[string]string {
	return map[string]string{"from": "which file/dir should be copied?", "to": "file/dir name to which copy"}
}

func copyFile(source string, dest string) (err error) {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err == nil {
		sourceInfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceInfo.Mode())
		}

	}

	return
}

func copyDir(source string, dest string) (err error) {
	sourceInfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, sourceInfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)
	objects, err := directory.Readdir(-1)

	for _, obj := range objects {
		sourceFile := source + "/" + obj.Name()
		destFile := dest + "/" + obj.Name()

		if obj.IsDir() {
			err = copyDir(sourceFile, destFile)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err = copyFile(sourceFile, destFile)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return
}
