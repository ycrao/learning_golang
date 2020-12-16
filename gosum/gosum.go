package main

/*
#include <stdlib.h>
*/
import "C"

//export Sum
func Sum(a int32, b int32) int32 {
	return a + b
}

// go build -ldflags "-s -w" -buildmode=c-shared -o gosum.dll
func main() {
}