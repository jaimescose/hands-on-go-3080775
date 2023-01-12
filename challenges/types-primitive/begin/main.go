// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"

func main() {
	// create a local variable "val" and assign it an int value
	val := 2

	// print the value of the local variable "val"
	fmt.Printf("%v=%T\n", val, val)

	// print the value of the package-level variable "val"
	printGlobal()

	// update the package-level variable "val" and print it again
	updateGlobal()
	printGlobal()

	// print the pointer address of the local variable "val"
	var pointer *int = &val
	fmt.Printf("%v (%T)\n", pointer, pointer)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*(pointer) = 100
	fmt.Printf("val = %v\n", val)
}

func printGlobal() {
	fmt.Printf("%v=%T\n", val, val)
}

func updateGlobal() {
	val = "new global val"
}
