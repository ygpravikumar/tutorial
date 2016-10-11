package main

import (
	"fmt"
	"strings"
)

//WordCount to return the wourdcount map in a string
func WordCount(s string) map[string]int {
	var wcmap = make(map[string]int)
	var ss = strings.Fields(s)
	for _, v := range ss {
		elem, ok := wcmap[string(v)]
		if ok == false {
			wcmap[string(v)] = 1
		} else {
			wcmap[string(v)] = elem + 1
		}
	}
	return wcmap
}

func main() {
	fmt.Println(WordCount("Hi there how are you?"))
}
