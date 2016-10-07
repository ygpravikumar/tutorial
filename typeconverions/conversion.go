package main

import "fmt"

// T(v) converts a variable v to Type T
// In go type conversions should be done explicitly
func main() {
	i := 42
	j := float64(i)
	k := uint(j)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}
