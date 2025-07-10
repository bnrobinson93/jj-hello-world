package main

import (
	"fmt"
)

// This is the best way to print things out
func main() {
	printHello()
	printGoodbye()
}

func printHello() {
	fmt.Println("Hello, world!")
}

func printGoodbye() {
	fmt.Println("Goodbye, world!")
}
