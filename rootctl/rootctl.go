package main

import (
	"fmt"
	"os"

	"github.com/TheDevtop/rootve/rootctl/cmdList"
	"github.com/TheDevtop/rootve/rootctl/cmdShell"
)

func usage() {
	fmt.Println("Usage: rootctl [command] [options]")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	cmdTag := os.Args[1]
	os.Args = os.Args[1:]

	switch cmdTag {
	case cmdList.TagLs:
		cmdList.LsMain()
	case cmdList.TagPs:
		cmdList.PsMain()
	case cmdShell.TagShell:
		cmdShell.ShellMain()
	default:
		usage()
		os.Exit(1)
	}

}
