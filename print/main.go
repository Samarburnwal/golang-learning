package main

import "fmt"

func main() {
	name := "samar"
	age := 45
	height := 5.5678890

	fmt.Println(name, age, height)
	fmt.Printf("name is %s, Age is %d, height is %.2f\n", name, age, height)
	fmt.Printf("Type of name is %T, age is %T, height is %T", name, age, height)

	//println automatically adds new line character, and spaces in between arguments
}
