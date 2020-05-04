package main

import "fmt"

type tytype int

var x tytype

func main() {

	fmt.Printf("x: %v %T\n", x, x)
	x = 42
	fmt.Printf("x: %v\n", x)
}
