package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/TheDevtop/rootve/pkg/libve"
)

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
		flagName     = flag.String("n", libve.NoneStr, "Specify virtual environment")
		flagOverride = flag.String("c", libve.NoneStr, "Specify command override")
	)
	flag.Usage = usage
	flag.Parse()

	// Load configuration, and find virtual environment
	// Warning: This section may panic in case of errors
	if mvc, err = libve.ReadConfig(libve.DefaultPath); err != nil {
		panic(err)
	}
	if vc, avail = mvc[*flagName]; !avail {
		panic("Virtual environment not found")
	}

	// Check if we need to override the command
	if *flagOverride != libve.NoneStr {
		vc.CommandPath, vc.CommandArgs = parseCommand(*flagOverride)
	}

	// Allocate virtual environment
	ve = libve.NewEnvironment(vc)

	// Print system name, and chroot
	fmt.Println(uname())
	if err = ve.Chroot(); err != nil {
		panic(err)
	}

	// Attach std devices
	ve.Attach(os.Stdin, os.Stdout, os.Stderr)

	// Become specified user/group
	if err = ve.SetCreds(); err != nil {
		panic(err)
	}

	// Mount filesystems, if possible
	if err = autoMount(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	// Execute the process, and finish
	if err = ve.Execute(); err != nil {
		fmt.Println(err)
	}
	os.Exit(0)
}