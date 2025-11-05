package main

import "fmt"

func main() {
	a := 2
	b := 4
	c := 5

	if a > 10 {
		fmt.Println("a is greater than 10")
	} else if a > 5 {
		fmt.Println("a is greater than 5")
	} else {
		fmt.Println("A is lesser than equal to 5")
	}

	if a > 10 && (b < 10 || c > 9) {
		fmt.Println("JSR")
	}
}
