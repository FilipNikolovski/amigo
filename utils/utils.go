package utils

import "strings"

func IsCommand(arg string) bool {
	return len(arg) >= 2 && arg[0] == '-' && arg[1] != '-'
}

func IsOpt(arg string) bool {
	return strings.HasPrefix(arg, "--")
}
