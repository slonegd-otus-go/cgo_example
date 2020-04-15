package main

/*
#include <stdint.h>

void inc(int size, int* res) {
	for (int i = 0; i < size; i++) {
		(*(res+i))++;
	}
}
*/
import "C"
import (
	"fmt"
	"unsafe"
	"reflect"
) 

func main() {
	c := []int{1,2,3,4}
	header := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	fmt.Printf("%+v\n",header)
	C.inc(C.int(len(c)), (*C.int)(unsafe.Pointer(header.Data)))
	fmt.Println(c)
}

// go run main.go 
// &{Data:824633811712 Len:4 Cap:4}
// [4294967298 4294967299 3 4]