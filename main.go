package main

/*
#include <stdint.h>

typedef struct {
	int a;
	int b;
	union {
		uint8_t c;
		uint16_t d;
	} u;
} some;

some v1 = {1,2,3};
some v2 = {1,2,{.d=4}};
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("%+v\n", C.v1) // {a:1 b:2 u:[3 0] _:[0 0]}
	fmt.Printf("%+v\n", C.v2) // {a:1 b:2 u:[4 0] _:[0 0]}
	fmt.Printf("%+v\n", getC(C.v1)) // 3
	fmt.Printf("%+v\n", getC(C.v1)) // 3
	fmt.Printf("%+v\n", getD(C.v2)) // 4
	fmt.Printf("%+v\n", getD(C.v2)) // 4
}

func getC(v C.some) uint8 {
	u := *(*C.uint8_t)(unsafe.Pointer(&v.u[0]))
	return uint8(u)
}

func getD(v C.some) uint16 {
	u := *(*C.uint16_t)(unsafe.Pointer(&v.u[0]))
	return uint16(u)
}
