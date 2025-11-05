package main

import "fmt"

func modifyNum(ptr *int){
	*ptr = 20
}

func main() {
	// var num int
	// num = 2

	// var ptr *int
	// ptr = &num

	num := 2
	num1 := "sdfg"
	ptr := &num1

	fmt.Println("Val of num is :", num)
	fmt.Println("Val of pointer is", ptr)

	fmt.Println("Val of num1 through pointer ptr is", *ptr)

	var pointer *int

	if pointer == nil {
		fmt.Println("pointer value is nil")
	}

	num2 := 2
	modifyNum(&num2)

	fmt.Println("Val of num2 is ",num2);
}
