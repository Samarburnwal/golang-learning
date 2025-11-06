package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Starting of the program")

	c := add(3, 4)
	defer fmt.Println("Value of addition is", c)
	defer fmt.Println("Middle of the program")
	fmt.Println("End of the program")

}

// the defer keyword makes the instruction execute at last of program and stores instructions in stack
// at last the instruction will execute in reverse order
