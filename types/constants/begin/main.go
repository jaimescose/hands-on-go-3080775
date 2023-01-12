// types/constants/begin/main.go
package main

import (
	"fmt"
	"unicode"
)

// declare a constant representing pi
const pi = 3.14159

// declare a rune constant with explicit type for the letter 'a'
// rune is an alias for the int32 type
const a rune = 'a'

func main() {
	// print the value and types of pi and a
	fmt.Printf("pi: %v - %T\n", pi, pi)
	// TODO: what is %c?
	fmt.Printf("a: %c - %T\n", a, a)

	// use unicode package to confirm that the rune is a letter
	fmt.Println(unicode.IsLetter(a))
}
