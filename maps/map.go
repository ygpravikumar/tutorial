package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

//map maps keys to values
//zero val of a map is nill
//length of a eero map is zero no key or values can be added to it
//make function initialises and returns a map ready for use
func main() {
	var mp map[string]int
	mp = make(map[string]int)
	mp["a"] = 1
	fmt.Println(mp["a"])

	//Map literals are like struct literals, but the keys are required.

	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{21, 45},
	}

	//If the top-level type is just a type name, you can omit it from the elements of the literal
	var gm = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
	fmt.Println(gm)

	vm := make(map[string]int)

	vm["Answer"] = 42
	fmt.Println("The value:", vm["Answer"])

	vm["Answer"] = 48
	fmt.Println("The value:", vm["Answer"])

	delete(vm, "Answer")
	fmt.Println("The value:", vm["Answer"])

	//Test that a key is present with a two-value assignment
	//If key is in m, ok is true. If not, ok is false.
	//If key is not in the map, then elem is the zero value for the map's element type.
	v, ok := vm["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
