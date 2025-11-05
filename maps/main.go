package main

import "fmt"

func main() {
	marks := make(map[string]int, 0)

	marks["samar"] = 99
	marks["rishav"] = 100
	marks["aman"] = 99

	for index, val := range marks {
		fmt.Printf("Marks of %s is %d\n", index, val)
	}

	// checking the value existed previously or not
	val, exist := marks["akash"]

	fmt.Println(val, ", exist ? :", exist)
	val1, exist1 := marks["akash"]
	fmt.Println(val1, ", exist ? :", exist1)

	person := map[string]int {
		"Alice" : 1,
		"Bbb" : 3,
		"Charlie" : 4,
	}

	for index, val := range person {
		fmt.Println(index, ":", val)
	}
}
