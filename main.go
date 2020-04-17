package main

/*
#include <stdint.h>

typedef struct {
	int a;
	int b;
	int c :2;
	int d :3;
} some;

some v1 = {1,2,3,4};
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("%+v\n", C.v1) // {a:1 b:2 _:[19 0 0 0]}
	fmt.Printf("%+v\n", getC(C.v1)) // 3
	fmt.Printf("%+v\n", getD(C.v1)) // 4
}

func getC(v C.some) int {
	p := uintptr(unsafe.Pointer(&v.b)) + unsafe.Sizeof(v.b)
	bf := *(*C.int)(unsafe.Pointer(p))
	return int(bf&3) // 0b11
}

func getD(v C.some) int {
	p := uintptr(unsafe.Pointer(&v.b)) + unsafe.Sizeof(v.b)
	bf := *(*C.int)(unsafe.Pointer(p))
	return int(bf>>2&7) // 0b111
}

