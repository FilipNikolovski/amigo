package cmd

import (
	"errors"
	"fmt"
)

type Command interface {
	Run() error
	Usage()
}

type arguments struct {
	cmd  string
	opts []string
	pkgs []string
}

func (a arguments) optExist(opts ...string) bool {
	for _, opt := range opts {
		if a.cmd[2:] == opt {
			return true
		}

		for _, o := range a.opts {
			if o == opt {
				return true
			}
		}
	}

	return false
}

func MakeCmd(cmd string, opts, pkgs []string) (Command, error) {
	switch cmd[0:2] {
	case "-S":
		return NewSyncCmd(cmd, opts, pkgs), nil
	case "-R":
		fmt.Println("Remove")
	case "-G":
		fmt.Println("Get")
	}

	return nil, errors.New("command not found")
}
