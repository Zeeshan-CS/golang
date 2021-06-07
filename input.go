package main

import (
	"bufio"
	"fmt"
	"os"
)

// for string
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type Something ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("You Typed: %q", input)

}
