// package main

// import "fmt"

// func main() {
// 	var i bool
// 	n := true
// 	j := 42
// 	fmt.Printf("%v , %T\n", n, n)
// 	fmt.Printf("%v , %T\n", j, j)
// 	fmt.Printf("%v , %T", i, i)
// 	a := 10
// 	b := 5

// 	fmt.Println("sum = ", a+b)
// 	fmt.Println("sub = ", a-b)
// 	fmt.Println("div = ", a/b)
// 	fmt.Println("mod = ", a%b)
// 	fmt.Println("mul = ", a*b)
// }

// package main

// // variable shadowing
// import (
// 	"fmt"
// )

// func test() (x int64, y int64) {
// 	x = 3
// 	y = 6

// 	return
// }
// func main() {
// 	var x int64 = 2
// 	fmt.Printf("x %v \n", x)

// 	x, y := test()
// 	if x > 0 {
// 		fmt.Printf("x %v y %v \n", x, y)
// 	}

// 	fmt.Printf("x %v \n", x)
// }

package main

type sum struct {
	X int
	Y int
}

func main() {

	// var i int
	// var j int

	// fmt.Println("Enter value of x")
	// fmt.Scanln(&i)

	// fmt.Println("Enter value of x")
	// fmt.Scanln(&j)

	// fmt.Println("Printing struct values", sum{i, j})

	// i := 43
	// j := &i
	// fmt.Println(i)
	// fmt.Println(j)

	//for continued
	// sum := 1
	// for sum < 1000 {
	// 	sum += sum + 2
	// }
	// fmt.Println(sum)

	// i := 0
	// sum := 0

	// for i = 1; i < 100; i += 2 {
	// 	sum += i
	// }
	// fmt.Println(sum)

}
