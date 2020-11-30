package gmod

import (
	"errors"
	"fmt"
	"github.com/fzft/gbuilder/pkg/cmd"
	"log"
	"strconv"
	"strings"
)

const GO_MOD_VERSION = 11

var ErrGoModVersion = errors.New(fmt.Sprintf("Go mod version required at least %v", GO_MOD_VERSION))

type GoVersion struct {
	version Version
}

type Version struct {
	major int
	minor int
}

func GOVNew() (*GoVersion, error) {
	v := &GoVersion{version: Version{}}
	out, err := cmd.NewCmd("go", "version").Run()
	if err != nil {
		log.Fatalf("failed to get go version\n")
		return v, err
	}

	vStr := strings.Split(strings.Split(out, " ")[2], ".")
	major, err := strconv.Atoi(vStr[1])
	if err != nil {
		log.Fatalf("failed to parse major version, %v\n", err)
		return v, err
	}
	minor, err := strconv.Atoi(vStr[2])
	if err != nil {
		log.Fatalf("failed to parse minor version, %v\n", err)
		return v, err
	}

	v.version.major = major
	v.version.minor = minor
	return v, nil
}

func (gv *GoVersion) GoMod(projpath string) error {
	_, err := cmd.NewCmd("cd", projpath).Run()
	if err != nil {
		log.Fatalf("failed to cd to Path %v\n", err)
		return err
	}
	_, err = cmd.NewCmd("go", "mod", "init").Run()
	if err != nil {
		log.Fatalf("failed to run go mod init %v\n", err)
		return err
	}
	return nil
}

func (gV *GoVersion) modValidate() error {
	if gV.version.major < GO_MOD_VERSION {
		return ErrGoModVersion
	}
	return nil
}

func (gv *GoVersion) Validate() error {
	if err := gv.modValidate(); err != nil {
		log.Fatalf("go mod validate fail %v\n", err)
		return err
	}
	return nil
}
