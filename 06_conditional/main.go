package main

import "fmt" //Format Package

func main() {
	// if else statment
	x := 10
	y := 10

	if x <= y {
		fmt.Printf("%d is less than or equal %d\n",x,y)
	}else{
		fmt.Printf("%d is less than %d\n",y , x)

	}
	// else_if statment
	color := "blue"
	if color == "red" {
		fmt.Printf("color is %s" , color)
	} else if color == "blue" {
		fmt.Printf("color is %s" , color)
	}else{
		fmt.Printf("color is %s" , color)
	}
      fmt.Print("\n")
	//logic operators
	var flag bool = false
	if flag && x == 10 {
		fmt.Println("flag and x are true condetions")
	}
	if flag || x == 10 {
		fmt.Println("flag or x is true condetions")
	}
	//switch
	switch color{
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	}
	//select for communication clause
}