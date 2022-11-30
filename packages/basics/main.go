// packages/basics/main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, Go!")

	fmt.Println("Today is", time.Now().Weekday())
}
