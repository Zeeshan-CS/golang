package main

import (
	"fmt"
)

func even_odd(x int) int {
	if x == 0 {
		fmt.Println("Please enter no greater than zero")
		return 3
	}

	if x%2 == 0 {
		return 0
	} else {
		return 1
	}
}

func main() {
	var num, x int

	fmt.Println("Enter a num to check its even or odd")
	fmt.Scanln(&num)

	x = even_odd(num)

	if x == 1 {
		fmt.Println("Odd Number")
	} else if x == 3 {
		fmt.Println("zero is a neutral num")
	} else {
		fmt.Println("Even Number")
	}
}
