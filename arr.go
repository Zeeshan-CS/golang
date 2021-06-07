package main

import "fmt"

func main() {
	fmt.Println("Enter value of num 1")
	var x int
	fmt.Scanln(&x)

	fmt.Println("Enter value of num 2")
	var y int
	fmt.Scanln(&y)

	fmt.Println("Enter 1 for add")
	fmt.Println("Enter 2 for div")
	fmt.Println("Enter 3 for sub")
	fmt.Println("Enter 4 for mul")
	var choice int
	fmt.Scanln(&choice)

	if choice == 1 {
		fmt.Println("sum of num 1 and num2 = ", x+y)
	} else if choice == 2 {
		fmt.Println("div of num 1 and num2 = ", x/y)
	} else if choice == 3 {
		fmt.Println("sub of num 1 and num2 = ", x-y)
	} else if choice == 4 {
		fmt.Println("mul of num 1 and num2 = ", x*y)
	}
}
