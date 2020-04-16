package main

/*
#include <stdint.h>

typedef struct {
	int a;
	int b;
} some;

void inc(some* v) {
	v->a++;
	v->b++;
}
*/
import "C"
import (
	"fmt"
) 

func main() {
	v := C.some{a:1,b:2}
	C.inc(&v)
	fmt.Printf("%+v\n", v) // {a:2 b:3}
}

