package main

import "fmt"

//[n]T defines an array of n values of type T
//Aray size if part of the datatype so arrays cannot be resized once they are declared
func main() {
	var a [2]string
	a[0] = "Hurray"
	a[1] = " World"
	fmt.Println(a)

	b := [5]int{1, 2, 3, 4, 5}
	var c = [3]string{"Ola", "Today", "typos"}
	var d = [5]int{1, 2, 3, 4, 5}
	fmt.Println(b, c, d)
}
