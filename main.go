package main

import (
	"fmt"
	"log"
	"os"
)

// func Start(c *exec.Cmd) error {
func main() {
	// c := exec.Command("/bin/bash")
	// Get pseudoterminal master from /dev/ptmx
	m, err := os.OpenFile("/dev/ptmx", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	fmt.Printf("%v\n", m)

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
