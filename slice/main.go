package main

import "fmt"

func main() {
	name := []string{}

	name = append(name, "samar")
	name = append(name, "rishav")
	fmt.Println("name :", name)
	fmt.Println("capacity :",cap(name))
	fmt.Println("length :",len(name))


	arr := make([]int, 3, 5)
	fmt.Println("name :", arr)
	fmt.Println("capacity :",cap(arr))
	fmt.Println("length :",len(arr))
	fmt.Println()

	arr = append(arr, 3, 5)
	fmt.Println("name :", arr)
	fmt.Println("capacity :",cap(arr))
	fmt.Println("length :",len(arr))
	fmt.Println()

	arr = append(arr, 6)
	fmt.Println("name :", arr)
	fmt.Println("capacity :",cap(arr))
	fmt.Println("length :",len(arr))
	fmt.Println()

	fmt.Printf("Type of name is : %T", name)
}
