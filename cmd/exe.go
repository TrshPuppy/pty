package cmd

import (
	"fmt"
	"os"
	"os/exec"

	// local mods:
	"pty/cmd/cgo"
)

func Start(cmd *exec.Cmd) error {

	// Get pseudoterminal master from /dev/ptmx
	m, err := os.OpenFile("/dev/ptmx", os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	// Go find the slave device:
	s := cgo.GetPTSName(m)

	defer m.Close()

	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", m)
	return nil
}
