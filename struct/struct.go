package main

import "fmt"

//A struct is a collection of fields
//Struct fields are accessed using dot notation

//Vertex is a collection of two integers
type Vertex struct {
	x, y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	fmt.Println(Vertex{2, 3}.y)

	// struct fields can be accessed through a struct pointer
	v := Vertex{1, 4}
	pv := &v
	// To dereference a strcut through its pointer we can use the dot notation just like on the struct itself
	// This is to avoid the cumbersome explicit dereference
	(*pv).x = 2
	pv.y = 10

	//A struct literal denotes a newly allocated struct value by listing the values of its fields.
	//You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)
	//The special prefix & returns a pointer to the struct value.
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{x: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p  = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println(v1, p, v2, v3)

	fmt.Println(*pv)

}
