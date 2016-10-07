package main

import "fmt"

// If declared without explicity type the type is infered from the value on the right hand side
// but when RHS contains an untyped expression, the inference is based on precission
func main() {
	i := 42
	j := 3.123456789123456567878
	k := 1 + 2i
	fmt.Printf("the type of i is %T\n", i)
	fmt.Printf("the type of i is %T\n", j)
	fmt.Printf("the type of i is %T\n", k)
}
