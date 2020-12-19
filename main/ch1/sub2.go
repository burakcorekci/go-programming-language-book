package ch1

import (
	"fmt"
	"go-programming-language-book/main/common"
	"os"
	"strconv"
	"strings"
)

func Sub2() {
	var args = CollectAllArgs()
	PrintSlice(args)
	ManualSlicePrint(args)
	StringsJoinSlicePrint(args)
	common.PrintSeparator()
}

func CollectAllArgs() []string {
	return os.Args
}

func CollectExtraArgs() []string {
	return CollectAllArgs()[1:]
}

func PrintSlice(stringSlice []string) {
	printWithTitle("Print slice", stringSlice)
}

func ManualSlicePrint(stringSlice []string) {
	var result, separator string
	for idx, str := range stringSlice {
		result += separator + strconv.Itoa(idx) + " " + str
		separator = "\n"
	}
	printWithTitle("Manual Slice Print", result)
}

func StringsJoinSlicePrint(stringSlice []string) {
	printWithTitle("Strings Join Slice Print", strings.Join(stringSlice, " "))
}

// Package private functions start with small letters
func printWithTitle(title string, content ...interface{}) {
	fmt.Println(title)
	fmt.Println(content...)
	fmt.Println()
}
