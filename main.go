package main

import (
	"log"
	"os/exec"

	// local mods:
	"pty/cmd"
)

// func Start(c *exec.Cmd) error {
func main() {
	c := exec.Command("/bin/bash")

	// Given the *exec.Cmd, we should start and return the PTYs and PTYmx:
	err := cmd.Start(c)
	if err != nil {
		log.Fatalf("Error starting PTY: %v\n", err)
	}
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
