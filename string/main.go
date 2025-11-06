package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orange,banana"
	strArr := strings.Split(data, ",")
	fmt.Println(strArr)

	str := "one two three four two nine seven"
	count := strings.Count(str, "two")
	fmt.Println(count)

	str = "                 Samar Burnwal              "
	trimmed := strings.TrimSpace(str)
	fmt.Println(trimmed)

	str1 := "Hello"
	str2 := "World"

	str = strings.Join([]string{str1, str2, "!"}," ")
	fmt.Println(str)
}
