package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(mul(4, 2))
}
