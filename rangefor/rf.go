package main

import "fmt"

//Range can be used to iterate over slice or a map[type]
//when iterating two vales are returned
//the first is the index and second a copy of the value at the index
//To skip the index use the _, to skip the value omit itentirely
func main() {
	a := [5]int{1, 2, 3, 4, 5}

	for index, value := range a {
		fmt.Println(index, value)
	}

	// Using just the value
	for _, value := range a {
		fmt.Println(value)
	}

	//Using just the index
	for index := range a {
		fmt.Println(index)
	}
}
