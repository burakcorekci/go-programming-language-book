package ch1

import (
	"bufio"
	"fmt"
	"os"
)

func Sub3() {
	fmt.Print(" - Chapter 1.3\n\n")
	duplicates()
}

func duplicates() {
	fmt.Println("Duplicates Map")

	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		txt := input.Text()
		if txt == "break" {
			break
		}
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	fmt.Println()
}
