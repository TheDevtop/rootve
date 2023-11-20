package main

import (
	"strings"

	"golang.org/x/sys/unix"
)

// Convert utsname structure to string
func uname() string {
	var (
		buf = new(unix.Utsname)
		err error
	)
	if err = unix.Uname(buf); err != nil {
		panic(err)
	}
	return string((*buf).Version[:])
}

// Seperate command path from arguments
func parseCommand(str string) (string, []string) {
	if argBuf := strings.Split(str, " "); len(argBuf) < 2 {
		return argBuf[0], nil
	} else {
		return argBuf[0], argBuf[1:]
	}
}
