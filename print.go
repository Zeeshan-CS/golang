package main

import (
	"fmt"
)

func main() {
	//format specifiers
	x := 567
	// bases
	fmt.Printf("hi %d\n", x)
	fmt.Printf("hi %b\n", x)
	fmt.Printf("hi %X\n", 4543)
	//floating number
	fmt.Printf("hi %e\n", 4.543)
	fmt.Printf("hi %E\n", 4.543)
	fmt.Printf("hi %f\n", 434.5545645646443)
	fmt.Printf("hi %e\n", 434.5545645646443)
	fmt.Printf("hi %g\n", 434.5545645646443)
	// String
	fmt.Printf("hi %s\n", "Zeeshan here")
	fmt.Printf("hi %q\n", "Zeeshan here")
	// Precision
	fmt.Printf("Number: %10q check space\n ", "Capreg")
	fmt.Printf("Number: %-10q check space\n ", "Capreg")
	fmt.Printf("Number: %-8q\n", 65)
	fmt.Printf("Number: %8q\n", 65)
	fmt.Printf("Number: %.2f\n", 56.3334444)
	fmt.Printf("Number: %.f\n", 56.3334444)

	// padding
	fmt.Printf("Number: %07d\n", 56)

	var out string = fmt.Sprintf("Num : %09d is good exam", 45)
	fmt.Println(out)
}
