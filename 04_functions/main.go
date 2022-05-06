package main

import (
	"fmt" //Format Package
	"math"
)

func greeting(name string) string{
	return "Hello " + name
}
func sum(num1 , num2 int) int {
	return num1 + num2
}
// return multi value
func swap(s1 , s2 string) (string , string){
	return s2 , s1
}
//use refrence instead of value
func addFive(value *int){
	*value += 5
}
//function closure
func getSequence() func() int {
	i := 0;
	return func() int {
		i += 1;
		return i;
	}
}
// methods
/* define a circle */
type Circle struct {
	x,y,radius float64
   }
   
   /* define a method for circle */
   func(circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
   }

func main() {
	
	fmt.Println(greeting("Paiman"))
	fmt.Println(sum(1 , 5));
	fmt.Println(swap("Jhon" , "Hi"))
	var number = 10
	fmt.Printf("Number is %d\n",number)
	//& to send address instead of value
	addFive(&number);
	fmt.Printf("Number is %d\n",number)
	// function as a value
	getSqureRoot := func (x float64) float64{
		return math.Sqrt(x);
	}
	fmt.Printf("Root of 64 is %f\n",getSqureRoot(64))
	// function closure
	var nextValue = getSequence();
	println("Next value is = ")
	fmt.Println(nextValue())
	fmt.Println(nextValue())
	fmt.Println(nextValue())

	// methods
	circle := Circle{x:0, y:0, radius:5}
	fmt.Printf("Circle area: %f", circle.area())
}
/*
call functions:
1- call by value
2- call by refrence(pointer)


functions type:
1-function as value
2-function closure
3-methods
*/