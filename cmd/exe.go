package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func Start(cmd *exec.Cmd) error {

	// Get pseudoterminal master from /dev/ptmx
	m, err := os.OpenFile("/dev/ptmx", os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", m)
	return nil
}
