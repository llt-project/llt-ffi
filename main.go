package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: llt-ffi/target/debug/libllt_ffi.dylib
#include "llt-ffi/include/llt.h"
*/
import "C"
import "fmt"

func main() {
	res := C.add(2, 3)
	fmt.Println(res)
}
