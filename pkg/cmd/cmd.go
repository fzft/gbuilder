package cmd

import (
	"log"
	"os/exec"
)

type Cmd struct {
	name    string
	options []string
}

func NewCmd(name string, arg ...string) Cmd {
	return Cmd{name: name, options: arg}
}

func (cmd Cmd) Run() (string, error) {
	c := exec.Command(cmd.name, cmd.options...)
	out, err := c.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd run failed with %s\n", err)
		return "", err
	}
	return string(out), nil
}
