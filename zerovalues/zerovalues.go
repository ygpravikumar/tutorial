package main

import "fmt"

// Variables declared without an initial value will be given a zero value
// For int = 0
// For string ""
// For bool false
func main() {
	var i int
	var j string
	var k bool

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}
