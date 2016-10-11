package main

import (
	"fmt"
)

//A defer statement defers the execution of a function untillthe sorrounding function returns
//The deferred function calls arguments are evaluated immediately
//but the function call is not executed untill the sorrounding function returns
func main() {

	// Deferred calls are pushed onto a stack
	// When the sorrounding function returns they are called in an LIFO order
	defer fmt.Println("This is Ravi.")
	defer fmt.Println("World")
	defer fmt.Println("Hello")
}
