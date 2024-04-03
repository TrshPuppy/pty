package cgo

// #include <stdlib.h>
// #include <stdint.h>
// #include <stdio.h>
//
// char *getPTSn(uint64_t fd) {
//     size_t len = 1;
//     char *buf;
//     int err;
//     do {
//        buf = malloc(len);
//        if(!buf){
// 	          return NULL;
//        };
//
//        err = ptsname_r(fd, buf, len);
//        if(err != 0) {
//           free(buf);
//           len++;
//           continue;
//        }
//     } while(err);
//
//     return buf;
// }
import "C"
import (
	"os"
)

func GetPTSName(f *os.File) string {
	i := f.Fd()
	name := C.GoString(C.getPTSn(C.ulong(i)))

	return name
}
