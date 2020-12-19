package common

import "fmt"

func PrintSubSectionSeparator() {
	printSeparator("-", 3)
}

func PrintChapterSeparator() {
	printSeparator("+", 10)
}

func printSeparator(char string, numberOfDashes int) {
	for i := 0; i < numberOfDashes; i++ {
		fmt.Print(char)
	}
	fmt.Print("\n\n")
}
