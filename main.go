package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName:   "Jim",
		lastName:    "Jackson",
		contactInfo: contactInfo{email: "jim@gmail.com", zipCode: 94000},
	}

	jim.updateName("Jimmy")
	jim.print()

}

func (p person) print() {
	fmt.Println(p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
