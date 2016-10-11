package main

// Just like for, if switch can also have an init statement
// The variables declared in the init statement are visible in the switch block
// a case body breaks automatically unless it ends with fallthrough statement
import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	// Switch cases are executed top to botton stopping when a case succeeds
	switch i := 20; os {
	case "OS X":
		{
			fmt.Println("This is linux", i)
		}
	case "windows":
		{

			fmt.Println("The os is windows amigos")
		}
	default:
		{
			fmt.Println(os)
		}
	}

	//Switch without a condition is the same as switch true.
	//This is a good alternative for a very long if-then-else chains
	switch {
	case 1 < -1:
		fmt.Println("1 < -1")
	case 1 > 20:
		fmt.Println("1 > 20")
	default:
		fmt.Println("This is a switch without a condition")
	}

}
