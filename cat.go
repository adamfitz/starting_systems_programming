// cat reads each file specified on the command line and write its contents to standard output
// usage: cat <file1> [<file2>] ...]
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for _, file := range os.Args[1:] { // 1. read ech file specified on the command line
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "open %s: %v", file, err)
			os.Exit(1)
		}
		// performance note: its better to use io.Copy, but here we are illustrating the process
		defer f.Close()
		b, err := io.ReadAll(f) // 2. read it into memory
		if err != nil {
			fmt.Fprintf(os.Stderr, "read %s: %v", file, err)
			os.Exit(1)
		}
		os.Stdout.Write(b) // 3. write its contents to standard output
	}
}
