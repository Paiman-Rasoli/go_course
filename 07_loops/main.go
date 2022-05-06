package main

import "fmt"

//Format Package
func main() {

/*	
	i := 1
  for i <= 10 {
	  fmt.Println(i)
	  i++
  }
  */
  
  /*
  for i := 1; i <= 10; i++{
	fmt.Println(i)
  }
  */
  //FizzBuzz
  /*
  for i := 1; i <= 100; i++{
	  if i % 15 == 0 {
		  fmt.Println("FizzBuzz")
	  }else if i % 3 == 0{
      fmt.Println("Fizz")
    }else if i % 5 == 0 {
      fmt.Println("Buzz")
    }else{
      fmt.Println(i)

    }
  }
  */
  /*
  // Infinite Loop
  for {
    fmt.Println("i am infinite")
  }
  */
  // break and continue
  var a int = 10
  for a < 20 {
     fmt.Printf("value of a: %d\n", a);
     a++;
     if a > 15 {
        /* terminate the loop using break statement */
        break;
     }
  }
}