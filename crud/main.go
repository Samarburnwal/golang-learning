package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performPostRequest() {
	todo := Todo{
		UserID:    123,
		Title:     "Samar Burnwal",
		Completed: true,
	}

	jsonData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("Error creating json")
		return
	}

	jsonString := string(jsonData)

	//converting json data to reader
	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos"
	res, err := http.Post(myUrl, "application/json", jsonReader)

	if err != nil {
		fmt.Println("Error sending request : ", err)
		return
	}

	defer res.Body.Close()

	// data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response : ", res.Status)
}

func PerformUpdateRequest() {
	todo := Todo{
		UserID:    2345,
		Title:     "Jai shree ram",
		Completed: false,
	}

	jsonData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("Error creating json")
		return
	}

	jsonString := string(jsonData)

	//converting json data to reader
	jsonReader := strings.NewReader(jsonString)
	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodPut, myUrl, jsonReader)

	if err != nil {
		fmt.Println("error creating PUT request :", err)
		return
	}

	req.Header.Set("Content-type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error fetching request", err)
		return
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response :", string(data))

	fmt.Println("Status :", res.Status)
}


func performDeleteRequest(){
	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodDelete, myUrl, nil)

	if err != nil {
		fmt.Println("error creating PUT request :", err)
		return
	}

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error fetching request", err)
		return
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response :", string(data))

	fmt.Println("Status :", res.Status)
}

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		fmt.Println("Error getting :", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting Response", res.Status)
		return
	}

	// data, err := ioutil.ReadAll(res.Body)

	// if err != nil {
	// 	fmt.Println("Error reading :", err)
	// 	return
	// }

	// fmt.Println("Data :", string(data))

	// var todo Todo

	// err = json.NewDecoder(res.Body).Decode(&todo)

	// if err != nil {
	// 	fmt.Println("Error decoding", err)
	// 	return
	// }
	// fmt.Println("Todo:", todo)

	// fmt.Println("Title response:", todo.Title)
	// fmt.Println("Completed response:", todo.Completed)
	// performPostRequest()
	// PerformUpdateRequest()
	performDeleteRequest()
}
