package main

//Format Package

// const MAX int = 3
func main() {
	/*
	// value 5 store in x
	x := 5;
	//address of x stores in y
	y := &x;
	fmt.Println("x and y", x , y)
	fmt.Printf("%T\n" , y)
	// use * to read the value from a pointer
	fmt.Printf("value in address %x is %d\n" , y , *y)
	// change value in pointer
	*y = 10
	fmt.Println("value in address y is ", *y)
	*/
	
	/*
	Nil pointer
	Go compiler assign a Nil value to a pointer variable 
	in case you do not have exact address to be assigned.
	This is done at the time of variable declaration.
       A pointer that is assigned nil is called a nil pointer.

       The nil pointer is a constant with a value of zero defined in several 
	 standard libraries. Consider the following program −
	*/
	// var id *int;
	// fmt.Printf("THE VALue of ptr is %x\n",id)
	/*
	1. pointer in arrays
	*/
	// a := []int{10,100,200}
	// var i int
	// var ptr [MAX]*int;
   
	// for  i = 0; i < MAX; i++ {
	//    ptr[i] = &a[i] /* assign the address of integer. */
	// }
	// for  i = 0; i < MAX; i++ {
	//    fmt.Printf("Value of a[%d] = %d\n", i,*ptr[i] )
	// }
	/*
	2. pointer to pointers
	*/
	// var a int = 3000
	/* take the address of a */
	// var ptr *int = &a
	/* take the address of ptr using address of operator & */
	// var pptr **int = &ptr

   
	/* take the value using pptr */
	// fmt.Printf("Value of a = %d\n", a )
	// fmt.Printf("Value available at *ptr = %d\n", *ptr )
	// fmt.Printf("Value available at **pptr = %d\n", **pptr)
	/* 3. pointers to funtions */
}