package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Usage: rootctl [command] [options]")
}

func main() {
	var (
		cmdFlag  string
		exitCode int
	)

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	cmdFlag = os.Args[1]
	os.Args = os.Args[1:]

	switch cmdFlag {
	case cmdLs:
		exitCode = lsMain()
	case cmdPs:
		exitCode = psMain()
	case cmdShell:
		exitCode = shellMain()
	case cmdStart:
		exitCode = startMain()
	case cmdStop:
		exitCode = stopMain()
	case cmdPause:
		exitCode = pauseMain()
	case cmdResume:
		exitCode = resumeMain()
	case cmdRemove:
		exitCode = removeMain()
	default:
		usage()
		exitCode = 1
	}

	os.Exit(exitCode)
}
