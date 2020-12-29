package ch1

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Sub3() {
	fmt.Print(" - Chapter 1.3\n\n")
	duplicates()
}

func duplicates() {
	fmt.Println("Duplicates Map")
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		files = []string{"main/resources/default.txt"}
	}
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "duplicates: %v\n", err)
			continue
		}
		countLinesStreaming(f, counts)
		_ = f.Close()
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	fmt.Println()
}

// Read the file line by line to process
func countLinesStreaming(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		txt := input.Text()
		if txt == "break" {
			break
		}
		counts[txt]++
	}
}

// Read whole file at once to process
func countLinesBatch(f *os.File, counts map[string]int) {
	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "countLinesBatch: %v\n", err)
		return
	}
	for _, line := range strings.Split(string(data), "\n") {
		counts[line]++
	}
}
