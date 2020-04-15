package main

/*
#include <stdint.h>

void x2(int size, int* res) {
	for (int i = 0; i < size; i++) {
		*(res+i) *= 2;
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
	C.x2(C.int(len(c)), (*C.int)(unsafe.Pointer(header.Data)))
	fmt.Println(c)
}
