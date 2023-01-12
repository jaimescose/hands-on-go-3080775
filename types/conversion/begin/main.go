// types/conversion/begin/main.go
package main

import "fmt"

// Go will create a new allocation for a converted value
func main() {
	// declare float and convert it to an unsigned int
	a := 42.5
	b := uint(a)

	fmt.Printf("%v=%T, %v=%T\n", a, a, b, b)
}
