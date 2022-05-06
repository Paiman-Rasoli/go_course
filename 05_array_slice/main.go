package main

import "fmt"

//Format Package

func main() {
	/*
	1. declartion and assignment
	var arrFruites [2]string
	//assign values
	arrFruites[0] = "Apple"
	arrFruites[1] = "Orange"
	fmt.Println(arrFruites , arrFruites[1])
	//decleare and assign
	users := [2]string{"Jhon" , "Bob"}
	fmt.Println(users)
	//dynamic length
	cars := []string{"Bmw" , "Toyta" , "Benz"}
	fmt.Println(cars)
	*/
	// methods
	cars := []string{"Bmw" , "Toyta" , "Benz" , "Honda"}
	// length of array
	fmt.Println(len(cars))
	//range
	fmt.Println(cars[1:3]) // 3 is exclude!
	//push

}