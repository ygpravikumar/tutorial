package main

import "fmt"

// Go has pointers
func main() {

	// *T refers to a pointer to type T
	// the zero value of a pointer is nil
	var i *int
	fmt.Println(i)

	// & operator assigns the address of a variable to a pointer
	j := 20
	i = &j

	// the * operator accesses the value referred by pointer
	// this acessing is called dereferencing or indirecting
	// dereferencing value of a nil pointer rises a panic
	*i = 40

	// go has no pointer arithmetic
	fmt.Println(i, *i)
}
