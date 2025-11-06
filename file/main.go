package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		file, err := os.Create("example.txt")

		if err != nil {
			fmt.Println("Error in creating file", err)
			return
		}

		defer file.Close()

		content := "Hello world by Samar"

		bytes, errors := io.WriteString(file, content)

		if errors != nil {
			fmt.Println("Error while writing into file")
			return
		}

		fmt.Println("Created file successfully", bytes)
	*/

	/*file, err := os.Open("example.txt")
	if err != nil{
		fmt.Println("Error while opening of file: ", err)
	}

	defer file.Close()

	//creating buffer for reading the file
	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error while reading file", err)
		}

		// Process the read content
		fmt.Println(string(buffer[:n]))
	}
	*/

	//Read entire data
	//importing ioutil gives a warning that it is deprecated
	//because the readfile function is also provided by OS
	content, err := os.ReadFile("example.txt")

	if err != nil {
		fmt.Println("Error while reading file")
		return
	}
	data := string(content)
	fmt.Println(data)
}
