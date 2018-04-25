package main

/*
#cgo CFLAGS: -I../DyLib
#cgo LDFLAGS: -L. -lc
#include <sys/utsname.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

const SYS_NAMELEN = 256

type utsname struct {
	Sysname [SYS_NAMELEN]byte	/* [XSI] Name of OS */
	Nodename [SYS_NAMELEN]byte	/* [XSI] Name of this network node */
	Release [SYS_NAMELEN]byte	/* [XSI] Release level */
	Version [SYS_NAMELEN]byte	/* [XSI] Version level */
	Machine [SYS_NAMELEN]byte	/* [XSI] Hardware type */
}

func CToGoString(c []byte) string {
    n := -1
    for i, b := range c {
        if b == 0 {
            break
        }
        n = i
    }
    return string(c[:n+1])
}


func main() {
	un := utsname{}
	C.uname((*C.struct_utsname)(unsafe.Pointer(&un)))
	fmt.Printf("%s\n", CToGoString(un.Sysname[:]))
	fmt.Printf("%s\n", CToGoString(un.Machine[:]))
}