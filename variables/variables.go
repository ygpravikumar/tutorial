//var statement declares a list of variables
//can be used at function or package level
//similar to function arguments, datatype is last

package main

import "fmt"

var c, python, java bool

func main() {

	//Can include initializers, one per variables
	var i int = 2
	var j, z bool = true, false

	//if an initializer is present, the type can be omitted
	var k = "Hello World!"

	//inside a function, the short assignment  statement can be used
	m := 2345
	fmt.Println(i, c, python, java)
	fmt.Println(j, z)
	fmt.Println(k)
	fmt.Println(m)
}
