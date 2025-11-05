package main

import "fmt"

func main() {
	var arr1 [5]int
	fmt.Println("array is :", arr1)

	var arr = [5]int{1, 2, 3, 4, 5}
	arr[1] = 56

	fmt.Println(arr)

	var name[5]string

	name[1] = "Samar"

	fmt.Println(name)
	fmt.Println(len(name))
	fmt.Printf("%q", name)
}
