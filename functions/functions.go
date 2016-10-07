package main

import "fmt"

//Functions can take multiple inputs
//Functions can return any number of results
//Named returns are possible in functions

//Add to add two input integers
func Add(x int, y int) int {
	return x + y
}

//Sub to subtract two input integers
func Sub(x, y int) int {
	return x - y
}

//Swap function to explain multiple returns
func Swap(x, y string) (string, string) {
	return y, x
}

//MultBy10And20 to explain named exports with a naked return
func MultBy10And20(x int) (y, z int) {
	y = x * 10
	z = x * 20
	// This is called a naked return
	// this returns named return values
	// these can harm readability and should only be used in short functions
	return
}

func main() {
	fmt.Println("The sum of 2 + 3 is", Add(2, 3))
	fmt.Println("The difference between 6 and 4 is", Sub(6, 4))
	x, y := MultBy10And20(10)
	fmt.Println("10 multiplied by 10, is ", x)
	fmt.Println("10 multiplied by 20, is ", y)

}
