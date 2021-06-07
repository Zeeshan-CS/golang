package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// for converting String into integer
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type your born year:  ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You wil be %d years old at the end of 2021", 2021-input)

}
