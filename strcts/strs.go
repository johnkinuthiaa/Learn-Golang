package main

import (
	"fmt"
)

func strcts() {
	/*
		structs acts as objects
	*/
	type person struct {
		firstName string
		lastName  string
		age       int
		course    string
	}
	john := person{
		firstName: "john",
		lastName:  "kinuthia",
		age:       19,
		course:    "Diploma in computer science",
	}
	toString(john.firstName, john.lastName, john.age, john.course)

}
func toString(
	firstName string,
	lastName string,
	age int,
	course string) {
	fmt.Println(
		"name", firstName+" "+lastName,
		"\nAge:", age,
		"\ncourse:", course)

}
func main() {
	strcts()
}
