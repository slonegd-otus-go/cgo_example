package main

/*
void sum(int a, int b, int* res) {
	*res = a + b;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
) 

func main() {
	a:=2
	b:=3
	var c int
	C.sum(C.int(a), C.int(b), (*C.int)(unsafe.Pointer(&c)))
	fmt.Println(c)
}
