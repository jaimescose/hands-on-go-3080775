// types/variables/begin/main.go
package main

import "fmt"

// declare package-level variables of type int
var d, e, f int // this will initialize them with default == 0

// declare package-level variables of type bool and override the default values (also known as "zero")
var x, y, z = true, false, true // type could be infered from the value

func main() {
	// print ints
	fmt.Println("The values for d, e, f are:", d, e, f)

	// override the default value of a package-level variable
	d = 1_000_000
	fmt.Printf("New value for d: %d\n", d)

	// declare and initialize variables of similar names as booleans but of local scope
	x, y, z := 0, 1.25, "hello" // this syntax in only available for local scope
	fmt.Println("x, y, z:", x, y, z)

	// print the package-level variables x, y, and z
	printVars()
}

func printVars() {
	fmt.Println("x, y, z:", x, y, z)
}