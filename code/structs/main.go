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
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

// pointer reminder
// turn adress into value with *address
// turn value into address with &value

func (p *person) updateName(firstName string) {
	p.firstName = firstName
}

func (p person) print() {
	fmt.Println(p)
}
