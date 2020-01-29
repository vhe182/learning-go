package main

import (
	"flag"
	"fmt"
	"os"
	"puzzleutil"
)

func PrintHelp(cmd string) {
	fmt.Fprintf(os.Stderr, "This tool is used to calculate and demonstrate solutions "+
		"to the 2019 advent of code challenges\n")
	fmt.Fprintf(os.Stderr, "Usage: %s <path/to/input-file>\n", os.Args[0])
}

func handlePuzzle1(inputFilename string) {
	_, err := os.Open(inputFilename)
	if err != nil {
		fmt.Printf("Failed to open resource file: %s\n", inputFilename)
		return
	}
	fmt.Printf("Mass Module file: %s\n", inputFilename)

	var masses []int
	masses = puzzleutil.ReadMassModuleFile(inputFilename)
	if masses == nil {
		fmt.Fprintf(os.Stderr, "ReadMassModuleFile returned nil\n")
		return
	}
	var fuelCost int = puzzleutil.CalculateFuel(masses)
	fmt.Fprintf(os.Stdout,
		"Total Fuel cost for file:%s is: %d", inputFilename, fuelCost)
}

func handlePuzzle2(inputFilename string) {
	fmt.Println("Implementation in progress")
}

/*
For Puzzle 2A:
1. Rewrite main() so that it takes a flag.  The flag will signal which puzzle
the user is attempting to solve (which functions to call).
2. Write well-defined tests to ensure requirements are met
  a. Test cases provided:
    i. 1,0,0,0,99 becomes 2,0,0,0,99 (1 + 1 = 2).
    ii. 2,3,0,3,99 becomes 2,3,0,6,99 (3 * 2 = 6).
    iii. 2,4,4,5,99,0 becomes 2,4,4,5,99,9801 (99 * 99 = 9801).
    iv. 1,1,1,4,99,5,6,0,99 becomes 30,1,1,4,2,5,6,0,99.
3. Implement the solution for puzzle 2A in a new function in puzzleutil package.
  a. Opcode 1 = addition
    i. Example: 1 10 20 30 ; a[30] = a[10] + a[20]
  b. Opcode 2 = multiplication
    i. Example: 2 3 0 3 ; b[3] = b[3] * b[0]
*/

func main() {
	wordPtr := flag.String("p", "puzzle=1", "The puzzle number to solve.")
	filenamePtr := flag.String("i",
		"input=resources/puzzle1-input.txt",
		"The input required to solve the puzzle.")
	flag.Parse()
	fmt.Println("puzzle: ", *wordPtr)
	fmt.Println("input filename: ", *filenamePtr)

	clArgs := os.Args[1:]
	if len(clArgs) != 4 {
		PrintHelp(os.Args[0])
		return
	}

	switch *wordPtr {
	case "1":
		handlePuzzle1(*filenamePtr)
	case "2":
		handlePuzzle2(*filenamePtr)
	}
}
