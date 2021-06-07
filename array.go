package main

import (
	"fmt"
)

func main() {
	numbers := [4][4]int{}
	var even []int
	var odd []int
	//row := 2
	//col := 4

	fmt.Println("Enter data in your 2-D Array")
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			fmt.Scanf("%d ", &numbers[row][col])
			// fmt.Println(col, row, numbers[row][col])
			if numbers[row][col]%2 == 0 {
				even = append(even, numbers[row][col])
			} else {
				odd = append(odd, numbers[row][col])
			}
		}
	}

	fmt.Println("Printing 2D Arrays", numbers)
	fmt.Println("Print Even Array", even)
	fmt.Println("Print Odd Array", odd)

	//Printing 2D Arrays [[1 2 3 4] [5 6 7 8] [9 10 11 12] [13 14 15 16]]
	//Print Even Array [2 4 6 8 10 12 14 16]
	//Print Odd Array [1 3 5 7 9 11 13 15]

}
