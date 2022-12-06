// packages/basics/main.go
package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func greeting() string {
	return "Hello"
}

func greetingWithName(name string) string {
	return "Hello, " + name + "!"
}

// the following function pattern return is called "naked return"
func greetingWithNameAndAge(name string, age int) (greeting string) {
	greeting = "Hello, my name is " + name + " and I have " + strconv.Itoa(age) + " years old!"
	return
}

// multi-value returns: a common pattern is that the last value retorn is used for errors
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func main() {
	fmt.Println("Hello, Go!")

	fmt.Println("Today is", time.Now().Weekday())

	fmt.Println(greeting())

	fmt.Println(greetingWithName("Jaime"))

	fmt.Println(greetingWithNameAndAge("Jaime", 26))

	fmt.Println(divide(10, 5))

	fmt.Println(divide(10, 0))
}
