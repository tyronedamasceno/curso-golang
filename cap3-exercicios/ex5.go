package main

import "fmt"

type tytype int
type tytype2 tytype

var y tytype2

func main() {

	fmt.Printf("y: %v %T\n", y, y)
	y = 42
	fmt.Printf("y: %v\n", y)
	y := tytype(y)
	fmt.Printf("y: %v %T\n", y, y)
}
