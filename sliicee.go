package main

import "fmt"

func main() {
	var x [5]int = [5]int{1, 2, 3, 4, 5}
	var s []int = x[1:3]
	fmt.Printf("%d\n", x)

	fmt.Printf("Slice Elements %d\n", s)

	//to return length
	fmt.Printf("Slice Lenth %d\n", len(s))
	fmt.Printf("Slice Capacity %d\n", cap(s))

}
