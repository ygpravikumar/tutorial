package main

import "fmt"

//Constants are declared like variables but with const keyword
//Can be numeric, boolean, string, charecter
//short assignment cannot be used for Constants

const Pi = 3.1457

func main() {
	const i = 2
	const j = "hello"
	const k = false

	// const keyword can be used in factored form
	// the type is infered from the assignment and precision
	const (
		m = 1
		l = 1 << 46
		n = 'c'
	)

	// Numeric constants are high precision integers
	const big = float64(1 << 100)

	// int is max 64 bit size so it cannot hold very big numbers like big here

	fmt.Println(Pi)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Printf("Type of m is %T\n", m)
	fmt.Printf("Type of l is %T\n", l)
	fmt.Printf("Type of n is %T\n", n)
	fmt.Printf("Type of big is %T\n", big)
	fmt.Println(big)
}
