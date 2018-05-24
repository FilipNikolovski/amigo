package cmd

import (
	"errors"
	"fmt"
)

type Sync struct {
	arguments
}

func NewSyncCmd(cmd string, opts, pkgs []string) Command {
	return &Sync{
		arguments: arguments{
			cmd:  cmd,
			opts: opts,
			pkgs: pkgs,
		},
	}
}

func (s *Sync) Run() error {
	if s.optExist("s", "--search") {
		return s.handleSearch()
	}

	return errors.New("invalid command. See usage using --help.")
}

func (s *Sync) handleSearch() error {
	fmt.Println("Search command")
	return nil
}

func (s *Sync) Usage() {
	fmt.Println(`
	USAGE:
		amigo -S [options] [pkg] [pkg] [...]
	OPTIONS:
		-s --search		search for packages
		-i --info		show info for packages
		-p --pkgbuild	show PKGBUILD
		-y --refresh	refresh database
		-u --upgrade	upgrade installed packages
		-c --clear		clear the cache directory
		-a --aur
		-m --maintainer
		-l --local
	SEE ALSO:
		-Sh		Show sync options and usage
		-Gh		Show get options and usage
			`)
}
