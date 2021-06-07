package main

import (
	"fmt"
	"time"
)

// func main() {

// 	go sayHello()
// 	time.Sleep(100 * time.Millisecond)
// }

// func sayHello() {
// 	fmt.Println("Salam Pakistan")
// }
func main() {
	var msg = "Go Routine"
	go func() {

		fmt.Println(msg)
	}()
	msg = "Good Bye"
	time.Sleep(100 * time.Millisecond)
}
