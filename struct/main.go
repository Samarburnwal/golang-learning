package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int

}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	HouseNo int
	Area    string
	State   string
}

type Employee struct {
	PersonDetails Person
	PersonContact Contact
	PersonAddress Address
}

func main() {
	var guy Person
	guy.firstName = "Samar"
	guy.lastName = "Burnwal"
	guy.age = 23

	fmt.Println(guy)

	person1 := Person{
		firstName: "Samixel",
		lastName:  "Singh",
		age:       34,
	}

	fmt.Println(person1)

	person2 := new(Person)
	person2.firstName = "Samarat"
	person2.lastName = "Singhal"
	person2.age = 10

	// since its made using new keyword, therefore person2 is a pointer
	fmt.Println(person2)

	var employee1 Employee

	employee1.PersonDetails = Person{
		firstName: "Kailash",
		lastName: "Kher",
		age: 25,
	}

	employee1.PersonContact = Contact{
		Email: "kailash234@gmail.com",
		Phone: "9876543210",
	}

	employee1.PersonAddress = Address{
		HouseNo: 20,
		Area: "Chandni Chowk",
		State: "Delhi",
	}

	fmt.Println(employee1);
}
