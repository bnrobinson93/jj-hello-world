// A "Hello, world!" program written in Go
package main

import (
	"fmt"
)

// This is the best way to print things out
func main() {
	print("Hello, world!")
	print("Goodbye, world!")
}

// A function to print out a string
func print(str string) {
	fmt.Println(str)
}
