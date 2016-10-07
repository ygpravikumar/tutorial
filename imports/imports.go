package main

// This is a factored import statement
// This can also be written as
// import "fmt"
// import "math/rand"
import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println(rand.Seed)
	fmt.Printf("A random number is %g, hurray", math.Sqrt(5))
}
