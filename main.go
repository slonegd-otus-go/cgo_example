package main

/*
#include <stdint.h>

typedef struct {
	int a;
	int b;
} some;

some ar[2] = {{1,2}, {3,4}};

some* get_some() { return ar; }
int   len()      { return 2; }
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	vs := C.get_some()
	len := C.len()
	fmt.Printf("%+v\n", vs)        // &{a:1 b:2}
	fmt.Printf("%+v\n", len)       // 2
	inc(int(len), vs)
	fmt.Printf("%+v\n", vs)        // {a:2 b:3}
	fmt.Printf("%+v\n", get(1,vs)) // {a:4 b:5}
}

func inc(len int, vs *C.some) {
	for i := 0; i < len; i++ {
		v := get(i, vs)
		v.a++
		v.b++
	}
}

func get(i int, vs *C.some) *C.some {
	p := uintptr(unsafe.Pointer(vs))
	p += uintptr(i)*unsafe.Sizeof(*vs)
	return (*C.some)(unsafe.Pointer(p))
}
