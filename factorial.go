package main

import "fmt"

func main() {

	var n int

	fmt.Println("Enter number to find factorial")
	fmt.Scanln(&n)
	f := 1
	i := 1
	for i <= n {
		f = f * i
		i++
	}
	fmt.Println("Factorial =", f)

}
