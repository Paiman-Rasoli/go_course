package main

import "fmt"

//Format Package

func main() {
	//key and value types are strings
	emails := make(map[string]string)
	emails["Bob"] = "bob@gmail.com"
	emails["paiman"] = "paimanrasoli789@gmail.com"

	fmt.Println(emails)
	fmt.Printf("Bob email is = %s\n" , emails["Bob"])
	fmt.Println(len(emails))
	// delete from map
	delete(emails , "Bob")
	fmt.Println(emails)
	//declare and assing
	ids := map[int]string{97318 : "paiman" , 7319 : "Jhon"}
	fmt.Println(ids)
}