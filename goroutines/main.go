package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello world")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("say Hello function ended successfully")
}

func sayHi() {
	fmt.Println("Hii Samar")
	time.Sleep(1000 * time.Millisecond)
}

func main() {
	fmt.Println("Learning Goroutines")
	go sayHello()
	sayHi()

	time.Sleep(2000 * time.Millisecond)
}
