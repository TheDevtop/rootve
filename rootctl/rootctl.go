package main

import (
	"fmt"
	"os"

	"github.com/TheDevtop/rootve/rootctl/cmdList"
	"github.com/TheDevtop/rootve/rootctl/cmdShell"
	"github.com/TheDevtop/rootve/rootctl/cmdState"
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
	case cmdState.TagStart:
		cmdState.StartMain()
	case cmdState.TagStop:
		cmdState.StopMain()
	case cmdState.TagPause:
		cmdState.PauseMain()
	case cmdState.TagResume:
		cmdState.ResumeMain()
	default:
		usage()
		os.Exit(1)
	}

}
