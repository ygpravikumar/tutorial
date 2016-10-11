package main

import (
	"fmt"
	"math"
)

//Sqrt function to calculate root by usings newtons formula
func Sqrt(x float64) float64 {
	z := float64(1)
	d := float64(1)
	temp := z
	for d > 1e-6 {
		temp = z
		z = z - ((z*z - x) / (float64(2) * z))
		d = math.Abs(z - temp)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
