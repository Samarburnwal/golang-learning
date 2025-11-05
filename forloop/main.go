package main

import "fmt"

func main() {
	for i := 0; i < 1; i++ {
		fmt.Println("Value :", i)
	}

	numbers := []int{123, 22, 37, 44, 59}

	for index, val := range numbers {
		fmt.Println(index, ":", val)
	}

	str := "Samar"

	for index, val := range str {
		fmt.Printf("val = %c, index = %d\n", val, index)
	}
}
