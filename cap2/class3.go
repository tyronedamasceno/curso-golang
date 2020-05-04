package main

import "fmt"

func main() {
	// := (walrus, marmota) - operador curto de declaração
	// = operador de atribuição

	x := 10
	y := "Hello world"

	fmt.Printf("x: %v %T\n", x, x)
	fmt.Printf("y: %v %T\n", y, y)

	x, z := 3, 4
	fmt.Printf("x: %v %T\n", x, x)
	fmt.Printf("z: %v %T\n", z, z)
}
