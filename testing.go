package main

import "fmt"

func main() {
	//cant re-declare but can shadow them
	// all vaiables must be used

	// Visibility:
	//1.lowercase first letter for package scope
	//2. uppercase first letter to export
	//3. no private scope

	// to print a string
	fmt.Println("its capregsoft")
	// 3 ways to declare variable
	var i int
	i = 44 // use for loop

	j := 41

	var k int = 55 // if go doest know th type

	fmt.Println(i, j, k)

}
