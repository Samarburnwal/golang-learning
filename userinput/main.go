package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Hello,", name, "biradar")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	fmt.Println(err,name)
}
