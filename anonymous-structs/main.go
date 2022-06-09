package main

import (
	"fmt"
	"strings"
)

func main() {

	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       30,
	}

	fmt.Printf("%#v\n", diana)

	type Book struct {
		string
		float64
		bool
	}

	b1 := Book{"1984 by George Orwell", 10.2, false}
	fmt.Printf("%#v\n", b1)

	fmt.Println(b1.string)

	type Employee1 struct {
		name   string
		salary int
		bool
	}

	e := Employee1{"John", 40000, false}
	fmt.Printf("%#v\n", e)
	e.bool = true

	fmt.Println(strings.Repeat("#", 10))

	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	john := Employee{
		name:   "John Keller",
		salary: 3000,
		contactInfo: Contact{
			email:   "jkeller@company.com",
			address: "Street 20, London",
			phone:   042324234,
		},
	}

	fmt.Printf("%+v\n", john)

	fmt.Printf("Employee's salary: %d\n", john.salary)

	fmt.Printf("Employee's email:%s\n", john.contactInfo.email)

	john.contactInfo.email = "new_email@thecompany.com"
	fmt.Printf("Employee's new email address:%s\n", john.contactInfo.email)

}
