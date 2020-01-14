package main

import (
	"fmt"
	"os"
	"puzzleutil"
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

	var masses []int
	masses = puzzleutil.ReadMassModuleFile(clArgs[0])
	if masses == nil {
		fmt.Fprintf(os.Stderr, "ReadMassModuleFile returned nil\n")
		return
	}
	var fuelCost int = puzzleutil.CalculateFuel(masses)
	fmt.Fprintf(os.Stdout,
		"Total Fuel cost for file:%s is: %d", clArgs[0], fuelCost)
}
