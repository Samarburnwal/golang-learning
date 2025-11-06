package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		fmt.Println("Error getting response", err)
		return
	}

	defer res.Body.Close()

	fmt.Printf("Type of response %T\n", res)
	fmt.Println("data :", res.Body)

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error reading response", err)
		return
	}

	fmt.Println("response:", string(data))
}
