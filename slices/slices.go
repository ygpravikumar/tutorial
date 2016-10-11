package main

import "fmt"

//Slices are flexible dynamic view into an array
//Slices are reference to an array they do not hold any data
//Changing data in an slices changes the array and all other slices pointing to the data will also see the changes

func main() {

	a := [5]int{1, 2, 3, 4, 5}
	var b []int = a[0:5]
	fmt.Println(b)
	fmt.Printf("The type of b is %T", b)

	//A slice literal is like an array literal, but without length
	r := []bool{true, true, false}
	fmt.Println(r)

	// Slice defaults
	// 0 for lower bound
	// length of the arrray for higher bound
	// While slicing 0 and length are defaults for bounds and if omitted are assumed so
	s := a[:]

	s = s[1:4]
	fmt.Println(s)
	fmt.Printf("the type of s is %T\n", s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	//Slice has both length and capacity
	//Length is the number of elements in array and
	//capacity is the number of elements in underlying array COUNTING from the first element in the slice
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(a)

	//A NILL slice is with zero elements and no underlying array
	// Its value is nil and len = cap = 0
	//the zero value of a slice is nil
	var ns []int
	fmt.Println(ns, len(ns), cap(ns))
	if ns == nil {
		fmt.Println("nil")
	}

	//make is used to create dynamic arrays and slices
	//make allocates a zeroed array and returns a slice that points to that array
	da := make([]int, 5)
	fmt.Println(len(da), cap(da))
	//To specify capacity pass a third argument to a make
	dc := make([]int, 2, 5)
	fmt.Println(dc, len(dc), cap(dc))
}
