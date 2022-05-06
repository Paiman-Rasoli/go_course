package main

import (
	"01_hello/main/03_packages/strutil"
	"fmt"
	"math"
)

func main()  {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(strutil.Reverse("olleh"))
}