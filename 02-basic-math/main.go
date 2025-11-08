package main

import (
	"fmt"
	"os"
)

func main() {
	var sum float64
	var n int

	for {
		var val float64

		if _, err := fmt.Fscanln(os.Stdin, &val); err != nil {
			break
		}
		sum += val
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "No values")
		os.Exit(-1)
	}
	// ctrl + d: to signal to end the running program

	fmt.Println("The average is:", sum/float64(n))
}

// make this read the nums from a text file using `go run . < nums.txt`
// or cat nums.txt | go run .
// so the program reads the standard input, from wherever it came from
