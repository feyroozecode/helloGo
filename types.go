package main

import "fmt"

func gotTypes() {

	/* User specified types  */
	const a int32 = 14
	const b float32 = 21.45
	var c complex128 = 1 + 4i
	var d uint16 = 11

	/* Default types */
	n := 42               // int
	pi := 3.14            // float64
	x, y := true, false   // bool
	z := "Go is awesome " // string

	fmt.Printf("user-specified types:\n %T %T %T %T \n", a, b, c, d)
	fmt.Printf("default types:\n %T %T %T %T %T", n, pi, x, y, z)
}
