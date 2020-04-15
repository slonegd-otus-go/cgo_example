package main

/*
int sum(int a, int b) {
	return a+b;
}
*/
import "C"
import "fmt"

func main() {
	a:=2
	b:=3
	c := C.sum(C.int(a), C.int(b))
	fmt.Println(c)
}
