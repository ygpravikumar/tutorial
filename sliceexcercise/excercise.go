package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

//Pic returns a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers
func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for i := range s {
		for j := 0; j < dx; j++ {
			s[i] = append(s[i], uint8(i+j))
		}
	}
	fmt.Println(len(s))
	for i := range s {
		fmt.Println(len(s[i]))
	}
	return s
}

func main() {
	pic.Show(Pic)
}
