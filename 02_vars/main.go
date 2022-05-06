package main

import "fmt"

/*
data types classify in 4 categories:
1-boolean types => true / false
2-numeric types => integer / floating
3-String
4-derived => pointer / array / structure / union / function / slice / interface / map / channel
*/
/*
varible declartion:
1-static type declartion
   var age int32 = 32;
2-dynamic type declartion
    age := 32
*/
func main(){
	// declare varible using var
	var name = "Paiman"
	var age int32 = 21
	// dynamic
	size := 1.6
	fmt.Println(age , size)
      fmt.Printf("Type of name is %T",name)
}