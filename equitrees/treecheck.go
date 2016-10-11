package main

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var fds = strings.Fields(t.String())
	for i := range fds {
		str := strings.Replace(fds[i], "(", "", -1)
		str = strings.Replace(str, ")", "", -1)
		if len(str) > 0 {
			i, err := strconv.Atoi(str)
			if err == nil {
				ch <- i
			}
		}
	}
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	var mp = make(map[int]int)
	for i := range ch1 {
		ct, ok := mp[i]
		if ok == false {
			mp[i] = ct + 1
		} else {
			delete(mp, i)
		}
	}

	for i := range ch2 {
		ct, ok := mp[i]
		if ok == false {
			mp[i] = ct + 1
		} else {
			delete(mp, i)
		}
	}
	return len(mp) == 0
}

func main() {

	var tr1, tr2 = tree.New(1), tree.New(1)
	var tr3 = tree.New(2)

	fmt.Println(Same(tr1, tr2))
	fmt.Println(Same(tr1, tr3))
}
