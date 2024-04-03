package cgo

// #include <stdlib.h>
// #include <stdint.h>
//
// int grant(uint64_t fd) {
//     int grantStatus = grantpt(fd);
//
//     return grantStatus;
// }
import "C"
import (
	"fmt"
	"os"
)

func GrantPT(mfd *os.File) error {
	var err error
	ifd := mfd.Fd()

	success := C.grant(C.ulong(ifd))
	if success != 0 {
		err = fmt.Errorf("Error granting permissions using grantpt()")
	}

	return err
}
