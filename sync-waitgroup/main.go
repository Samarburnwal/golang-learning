package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Worker", i, "started")
	//Something work here
	fmt.Println("Worker", i, "ended")

}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		go worker(i, &wg)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println("Everything completed successfully")
}
