package cmd

import (
	"fmt"
	"os"
	"os/exec"

	// local mods:
	"pty/cmd/cgo"
)

func Start(c *exec.Cmd) error {

	// Get pseudoterminal master from /dev/ptmx
	mptr, err := os.OpenFile("/dev/ptmx", os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	defer mptr.Close()

	// Get the name of the slave device:
	sname, err := cgo.GetPTSName(mptr)
	if err != nil {
		return err
	}

	// Now that we have the name, we have to call grantpt() & unlockpt():
	err = cgo.GrantPT(mptr)
	if err != nil {
		return err
	}

	err = cgo.UnlockPt(mptr)
	if err != nil {
		return err
	}
	fmt.Printf("PTS name: %s\n", sname)
	// Now that permission is granted, and the slave is unlocked, we can open the pts device file:
	return nil
}
