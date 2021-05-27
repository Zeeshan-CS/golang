package main

import "fmt"

func main(){
	/* This is my first sample program. */
	var x , y ,	 z int
	
	fmt.Println("Enter Value of x")
	fmt.Scanln(&x)
		
	fmt.Println("Enter Value of y")
	fmt.Scanln(&y)
	
	z = x + y
	fmt.Println("Sum of x and y = " , z)

}