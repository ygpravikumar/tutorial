package main

import "fmt"

// Funcion are values they can be passed around like other values
// Function values my be used as arguments and return values

func Add(x, y int) int {
	return x + y
}

func All(fn func(x, y int) int) int {
	return fn(3, 5)
}

func main() {
	fmt.Println(All(Add))
}
