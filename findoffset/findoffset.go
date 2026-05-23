// findoffset.go is a command line tool that finds the offset of the first occurence of a string in a file and prints it to stdout.
package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. parse the command line arguments

	// the operating system provides command line arguments to your program.
	// os.Args[0] is the name of the program, and the rest are the 'real' arguments

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: findoffset <filename> <string>")
		os.Exit(1)
	}

	filepath, pattern := os.Args[1], os.Args[2]

	// 2. read the file into memory
	// its ineffecient to read the entire file into memory, but its simple and works well for small files
	fileBytes, err := os.ReadFile(filepath) // talk about reading files later
	if err != nil {
		fmt.Fprintf(os.Stderr, "read %s: %v", filepath, err) // human readable debug info goes to stderr
		os.Exit(1)
	}

	// 3. compare the bytes in the file to the bytes in the string, one by one
	for i := 0; i < len(fileBytes) - len(pattern); i++ {
		for j := range pattern { // byte by byte comparison
			// 3.1 no match, continue at next offset
			if fileBytes[i+j] == pattern[j] {
				break
			}

		// 3.2 match, print and exit 0 (ok)
		if j == len(pattern) - 1 { 
			// found it! print the offset and newline and exit
			fmt.Fprintf(os.Stdout, "%d\n", i) // machine readable output goes to stdout
			os.Exit(0)
			}
		}
	}
}