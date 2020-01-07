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
	fmt.Printf(clArgs[0])
	fmt.Println()
	if puzzleutil.ReadMassModuleFile(clArgs[0]) < 0 {
		fmt.Println("Error: Failed to read in mass module file")
		return
	}
	fmt.Println()
	fmt.Println(strconv.Itoa(puzzleutil.CalculateFuel(10)))
	fmt.Println()
	fmt.Printf(stringutil.Reverse(clArgs[0]))
}
