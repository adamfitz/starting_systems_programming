// echo prints its argument to standard output, separated by spaces and terminated by a newline
// usage: echo <args...>
package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. iterate over the command line arguments
	for i, arg := range os.Args[1:] {
		if i > 0 {
			fmt.Print(" ")
		}
		// 2. print each argument to standard output, separated by spaces
		fmt.Print(arg)
	}
	// 3. terminate with a newline
	fmt.Println()
}

