package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/TheDevtop/rootve/pkg/libve"
	"golang.org/x/sys/unix"
)

const noneStr = "[None]"

func uname() {
	var (
		buf = new(unix.Utsname)
		err error
	)
	if err = unix.Uname(buf); err != nil {
		panic(err)
	}
	fmt.Printf(
		"%s %s %s\n",
		string((*buf).Sysname[:]),
		string((*buf).Release[:]),
		string((*buf).Version[:]),
	)
}

func usage() {
	fmt.Println("rootexec: Create VE and execute process")
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	var (
		vc  libve.VirtConfig
		ve  *libve.VirtEnv
		err error
	)

	// Setup and parse flags
	var (
		optRoot    = flag.String("r", noneStr, "Specify chroot path")
		optDir     = flag.String("d", noneStr, "Specify working directory")
		optCommand = flag.String("c", noneStr, "Specify command")
		optClean   = flag.Bool("e", false, "Specify clean environment")
	)
	flag.Usage = usage
	flag.Parse()

	// The flags need to be specified explicitly
	if *optRoot == noneStr || *optDir == noneStr || *optCommand == noneStr {
		usage()
	}

	// Split the command option into path and argument
	if argBuf := strings.Split(*optCommand, " "); len(argBuf) < 2 {
		vc.CommandPath = argBuf[0]
		vc.CommandArgs = nil
	} else {
		vc.CommandPath = argBuf[0]
		vc.CommandArgs = argBuf[1:]
	}

	// Fill the configuration in
	vc.Root = *optRoot
	vc.Directory = *optDir
	vc.Clean = *optClean

	// Allocate virtual environment, and attach std devices
	ve = libve.NewEnvironment(vc)
	ve.Attach(os.Stdin, os.Stdout, os.Stderr)

	// Print system name, and chroot
	// Warning: This section may panic in case of errors
	uname()
	if err = ve.Chroot(); err != nil {
		panic(err)
	}

	// Execute the process, and finish
	if err = ve.Execute(); err != nil {
		panic(err)
	}
	os.Exit(0)
}
