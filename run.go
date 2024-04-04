package main

import (
	"log"
	"os"
	"os/exec"

	// local mods:
	"pty/cmd"
)

// func Start(c *exec.Cmd) error {
func GetPseudoterminalDevices() (*os.File, string) {
	c := exec.Command("/bin/bash")

	// Given the *exec.Cmd, we should start and return the PTYs and PTYmx:
	mDevice, sDevice, err := cmd.Start(c)
	if err != nil {
		log.Fatalf("Error starting PTY: %v\n", err)
	}

	return mDevice, sDevice
}

/*
	handle all the pty things
	- start /dev/ptmx
	- return slave
	- return master
	- handle subprocesses spawned from original bash process

	look into:
		- setting stdin/stdout/stderr
		- handle sizing
			- sizes dynamically?
				- ex: were on the other end of a socket




	two struct/ class things
		- slave
		- master
*/
