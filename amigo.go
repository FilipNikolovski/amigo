package main

import (
	"fmt"
	"os"

	"github.com/FilipNikolovski/amigo/cmd"
)

func usage() {
	fmt.Println(`NAME:
	amigo - An AUR helper
USAGE:
	amigo [option] [pkg] [pkg] [...]
COMMANDS:
	-S		Install package
	-G		Clone package in the current dir
	-R		Remove package
OPTIONS:
	--help		Show help
	--version	Show version
SEE ALSO:
	-Sh		Show sync options and usage
	-Gh		Show get options and usage
		`)
}

func version() {
	fmt.Println("0.1.0")
}

func parse(args []string) (cmd string, opts []string, pkgs []string, err error) {
	if len(args) < 2 {
		err = fmt.Errorf("Please specify a command to run")
		return
	}

	cmd = args[1]

	if len(cmd) == 1 || cmd == "-H" || cmd == "-h" || cmd == "--help" {
		usage()
		os.Exit(0)
	}

	if cmd == "-V" || cmd == "-v" || cmd == "--version" {
		version()
		os.Exit(0)
	}

	for _, arg := range args[2:] {
		if arg[0] == '-' {
			opts = append(opts, arg)
			continue
		}

		pkgs = append(pkgs, arg)
	}

	return
}

func main() {
	command, opts, pkgs, err := parse(os.Args)
	if err != nil {
		usage()
		os.Exit(0)
	}

	fmt.Printf("Command: %s Options: %+v Packages: %+v\n", command, opts, pkgs)
	c, err := cmd.MakeCmd(command, opts, pkgs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = c.Run()
	if err != nil {
		fmt.Println(err)
	}
}
