package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	// firstName string
	// lastName  string
	// location  string
	// age       int
	firstName , lastName ,location  string
	age int
}
// greeting (value reciver)
func (p Person) greeting() string{
	return "My full name is "+ p.firstName + " "+p.lastName + " and I am " +
	 strconv.Itoa(p.age)
}
//hasBirthday(pointer reciver)
func (p *Person) hasBirthday() {
	p.age++
}
//changeLastName(pointer reciver)
func (p *Person)changeLastName(newLastName string){
	p.lastName = newLastName
}
func main() {
	person1 := Person{firstName: "Paiman", lastName: "Rasoli", location: "Afghanistan", age: 21}
	//alternative
	// person1 := Person{ "Paiman",  "Rasoli",  "Afghanistan",  21}
	person1.hasBirthday()
	person1.changeLastName("Bob")
	fmt.Println(person1.greeting())
}