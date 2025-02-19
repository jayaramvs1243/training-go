package main

import "fmt"

func main() {
	// 1. Create and execute an anonymous function
	func() {
		fmt.Println("Creating and executing an anonymous function")
	}()

	// 2. Store an anonymous function in a variable
	var_anonymous_function := func(name string) {
		fmt.Printf("This is an anonymous function stored in a variable. Passed parameter: %s", name)
	}
	var_anonymous_function("Jayaram")
}
