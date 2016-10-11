package main

import (
	"fmt"
	"math"
)

// if statement is like for, the expression need not be sorrounded by braces, but
// the parenthesis is needed

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

// just like for if can start with an init statement
// Variables declared in the init statement are just in the scope of the if block andits correspoding else blocks
func max(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	} else {
		fmt.Printf("The power %g is more than limit %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(sqrt(4))
	fmt.Println(sqrt(-5))
	// Both calls to the max are executed and returns before the call to Println takesplace
	fmt.Println(max(2, 3, 18), max(3, 3, 18))
}
