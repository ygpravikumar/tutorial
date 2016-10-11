package main

import "fmt"

//A closure is a function value that references variables from outside its body.
//The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
//Each closure has its own copies of re variables
//adder returns a closure function
func adder() func(x int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

}
