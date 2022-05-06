/*
The range keyword is used in for loop to iterate over items of an array, slice, channel or map.
*/
package main

import "fmt"
func main(){
//slice
  ids := []int {10,20,30,40,50,60,70,80}
  for i := range ids {
	  fmt.Printf("index %d and value is %d\n" , i , ids[i])
  }
  //add ids
  sum := 0
  for _, id := range ids {
	  sum += id
  }
  fmt.Printf("Sum is = %d\n" , sum)
  // range with maps
  data := map[int]string{97318 : "paiman" , 7319 : "Jhon"}
  for k, v := range data{
	  fmt.Println("Key is ",k , "and value is ",v)
  }
}