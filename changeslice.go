package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 5}
	var s []int = a[3:5]
	fmt.Println(s)
	fmt.Println("Capacity of S %d ", cap(s))
	fmt.Println("length of S %d ", len(s))
	b := append(a, 20)
	fmt.Println("Capacity of b %d ", cap(b))
	fmt.Println("Length of b %d ", len(b))
	b = append(b, 20)
	b = append(b, 20)
	b = append(b, 20)
	b = append(b, 20)
	b = append(b, 20)

	fmt.Println("Capacity of b %d ", cap(b))
	fmt.Println("Length of b %d ", len(b))
	fmt.Println(b)

}
