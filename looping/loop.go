package main

import "fmt"

//Go has only one looping struct, for loop
func main() {
	sum := 0
	// For loop has 3 components seperated by semicolons
	// Init statement executed before the iteration
	// condition executed during each iteration
	// post statement executed after the end of each iteration
	// the loop will stop once the condition is false
	// there are no braces and the parenthesis is required
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//the init and post statements are optional
	//at this point we can even drop the semicolons
	for sum < 1000 {
		sum += sum
	}

	// if we omit the condition a for will be an infinite loop

	fmt.Println(sum)
}
