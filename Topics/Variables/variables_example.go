package main

import "fmt"

func main() {

	// Declaring single variable

	var a int = 1   // Declared with var
	const b int = 2 // Declared with const

	var c int = 3 // Initialized and Assigned with Type
	var d = 4     // Initialized and Assigned without Type

	e := 5 // Declared with short hand

	var f int // Initialized with type
	f = 6     // Assigned value to existing variable

	// Declaring Multiple variables

	var g, h, i int = 7, 8, 9         // Multiple variables of same type in one line
	var j, k, l = 10, " hello ", true // Multiple variables of different types in one line

	m, n := " a ", false // Multiple types of variables in single line shorthand

	var (
		o int  = 11     //
		p bool = true   // Multi type and multi line variable declaration
		q      = " hi " // inside bracket
		r int           //
	)

	r = 12 // Initialized inside brackets, Assigned outside

	// Prints all variables
	fmt.Print(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
}
