//Slice of Slice example
package main

import (
	"fmt"
	"strings"
)

func main() {
	//Slice can contain anything including other slices
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	//use the built in append function to append to a slices
	//if the underlying array is small a new bigger array will be allocated and the new slice will point to the new array
	a := [5]int{1, 2, 3, 4, 5}
	s := a[:]
	s = append(s, 6, 7, 8, 9)
	fmt.Println(s, a)
}
