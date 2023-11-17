package main

import (
	"fmt"
	"os/exec"
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

// Mount filesystems, if possible
func autoMount() error {
	cmd := exec.Command("/sbin/mount", "-a")
	if cmd.Run() != nil {
		return fmt.Errorf("could not mount all filesystems properly")
	}
	return nil
}
