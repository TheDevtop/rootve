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
		flagName  = flag.String("n", noneStr, "Specify virtual environment")
		flagShell = flag.Bool("s", false, "Specify interactive session")
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

	// Initialize networking
	if err = ve.Netinit(); err != nil {
		fmt.Println(err)
	}

	// Configure the standard/console devices
	if err = ve.Stdinit(*flagShell); err != nil {
		fmt.Println(err)
	}

	// If we want an interactive session,
	// we will not call the cleanup functions afterwards.
	// As other processes are still active
	if *flagShell {
		// Set the process to be /bin/ksh
		ve.SetShell()
		// Execute the process
		if err = ve.Execute(); err != nil {
			fmt.Println(err)
		}
		os.Exit(0)
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
