package main

import (
	"fmt"
	"os"

	"github.com/FilipNikolovski/amigo/utils"
)

func usage() {
	fmt.Println(`NAME:
	amigo - An AUR helper
USAGE:
	amigo [option] [pkg] [pkg] [...]
COMMANDS:
	-S		Install package
	-Ss		Search for package
OPTIONS:
	--help		Show help
	--version	Show version
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

	for _, arg := range args[1:] {
		if utils.IsCommand(arg) {
			switch arg {
			case "-H", "-h":
				arg = "--help"
			case "-V", "-v":
				arg = "--version"
			default:
				cmd = arg
				continue
			}
		}

		if utils.IsOpt(arg) {
			switch arg {
			case "--help":
				usage()
				os.Exit(0)
			case "--version":
				version()
				os.Exit(0)
			}

			continue
		}

		pkgs = append(pkgs, arg)
	}

	return
}

func main() {
	_, _, _, err := parse(os.Args)
	if err != nil {
		usage()
		os.Exit(0)
	}
}
