package main

import "fmt"

func main() {
	var next_mrk int
	fmt.Println("Enter Marks of 10 students")
	fmt.Println("Mrks Range 1 - 100")
	sum := 0
	avg := 0
	i := 0
	for i < 10 {
		fmt.Scanln(&next_mrk)
		for next_mrk < 1 || next_mrk > 100 {
			fmt.Println("You Enter marks out of range")
			fmt.Println("Re-Enter Mrks")
			fmt.Println("Range 1 to 100")

			fmt.Scanln(&next_mrk)
		}

		sum += next_mrk
		i++
	}
	avg = sum / 10
	fmt.Println("sum = ", sum)
	fmt.Println("Avg = ", avg)

}
