package main

import (
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
