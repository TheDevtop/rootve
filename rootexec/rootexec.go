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

func parseCommand(str string) (string, []string) {
	if argBuf := strings.Split(str, " "); len(argBuf) < 2 {
		return argBuf[0], nil
	} else {
		return argBuf[0], argBuf[1:]
	}
}

func usage() {
	fmt.Println("rootexec: Create VE and execute process")
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	var (
		mvc   map[string]libve.VirtConfig
		vc    libve.VirtConfig
		avail bool
		ve    *libve.VirtEnv
		err   error
	)

	// Setup and parse flags
	var (
		flagName     = flag.String("n", noneStr, "Specify virtual environment")
		flagOverride = flag.String("c", noneStr, "Specify command override")
	)
	flag.Usage = usage
	flag.Parse()

	// Load configuration, and find virtual environment
	if mvc, err = libve.ReadConfig(libve.DefaultPath); err != nil {
		panic(err)
	}
	if vc, avail = mvc[*flagName]; !avail {
		panic("Virtual environment not found")
	}

	// Check if we need to override the command
	if *flagOverride != noneStr {
		vc.CommandPath, vc.CommandArgs = parseCommand(*flagOverride)
	}

	// Allocate virtual environment, and attach std devices
	ve = libve.NewEnvironment(vc)
	ve.Attach(os.Stdin, os.Stdout, os.Stderr)

	// Print system name, and chroot
	// Warning: This section may panic in case of errors
	fmt.Println(uname())
	if err = ve.Chroot(); err != nil {
		panic(err)
	}

	// Become specified user/group
	if err = ve.SetCreds(); err != nil {
		panic(err)
	}

	// Execute the process, and finish
	if err = ve.Execute(); err != nil {
		fmt.Println(err)
	}
	os.Exit(0)
}
