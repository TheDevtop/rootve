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
	const noneStr = "[None]"

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
		flagAttach   = flag.Bool("a", false, "Specify attached execution")
		flagOverride = flag.String("c", noneStr, "Specify command override")
	)
	flag.Usage = usage
	flag.Parse()

	// Load configuration, and find virtual environment
	// Warning: This section may panic in case of errors
	if mvc, err = libve.ReadConfig(libve.ConfigPath); err != nil {
		panic(err)
	}
	if vc, avail = mvc[*flagName]; !avail {
		panic("Virtual environment not found")
	}

	// Check if we need to override the command
	if *flagOverride != noneStr {
		vc.CommandPath, vc.CommandArgs = parseCommand(*flagOverride)
	}

	// Allocate virtual environment
	ve = libve.NewEnvironment(vc)

	// Print system name, and chroot
	fmt.Println(uname())
	if err = ve.Chroot(); err != nil {
		panic(err)
	}

	// Become specified user/group
	if err = ve.SetCreds(); err != nil {
		panic(err)
	}

	// Mount filesystems
	if err = ve.Mount(); err != nil {
		fmt.Println(err)
	}

	// Initialize devices
	ve.Devinit()

	// Configure the standard/console devices
	if err = ve.Stdinit(*flagAttach); err != nil {
		fmt.Println(err)
	}

	// Initialize networking
	if err = ve.Netinit(); err != nil {
		fmt.Println(err)
	}

	// Execute the process
	if err = ve.Execute(); err != nil {
		fmt.Println(err)
	}

	// Unmount filesystems and exit
	if err = ve.Umount(); err != nil {
		fmt.Println(err)
	}
	os.Exit(0)
}
