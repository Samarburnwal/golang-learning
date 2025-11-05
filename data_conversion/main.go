package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 10
	str := strconv.Itoa(num) // I -> integer, to, a -> alphabet

	fmt.Printf("%T, %T\n", num, str)

	var data float64 = float64(num)
	data = data + 1.23

	fmt.Println("Data is", data)
	fmt.Printf("Type of data is %T\n", data)

	number_string := "1234"
	number_int, _ := strconv.Atoi(number_string)

	fmt.Printf("Type of number_int %T\n", number_int)
	float_string := "3.14"
	float_number, _ := strconv.ParseFloat(float_string, 64)
	fmt.Println("number float is", float_number)
	fmt.Printf("Type is %T\n", float_number)
}
