package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first := 0
	sec := 1
	count := 0
	return func() int {
		count++
		if count == 1 {
			return 0
		} else if count == 2 {
			return 1
		}
		temp := sec
		sec = sec + first

		first = temp
		return sec
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
