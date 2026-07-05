package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: ffi/target/debug/libffi.dylib
#include "ffi/include/llt.h"
*/
import "C"
import "fmt"

func main() {
	res := C.add(2, 3)
	fmt.Println(res)
}
