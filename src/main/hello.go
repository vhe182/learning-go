package main

import (
	"fmt"
	"os"
	"puzzleutil"
	"strconv"
	"stringutil"
)

func PrintHelp(cmd string) {
	fmt.Fprintf(os.Stderr, "This tool is used to calculate and demonstrate solutions "+
		"to the 2019 advent of code challenges\n")
	fmt.Fprintf(os.Stderr, "Usage: %s <path/to/input-file>\n", os.Args[0])
}

func main() {
	clArgs := os.Args[1:]
	if len(clArgs) != 1 {
		PrintHelp(os.Args[0])
		return
	}
	_, err := os.Open(clArgs[0])
	if err != nil {
		fmt.Printf("Failed to open resource file.\n")
		return
	}
	fmt.Printf("Mass Module file: %s\n", clArgs[0])

	var numbers []int
	numbers = puzzleutil.ReadMassModuleFile(clArgs[0])
	if numbers == nil {
		fmt.Fprintf(os.Stderr, "ReadMassModuleFile returned nil\n")
		return
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Value %d is %d\n", i, numbers[i])
	}
	fmt.Println()
	fmt.Println(strconv.Itoa(puzzleutil.CalculateFuel(10)))
	fmt.Println()
	fmt.Printf(stringutil.Reverse(clArgs[1]))
}
