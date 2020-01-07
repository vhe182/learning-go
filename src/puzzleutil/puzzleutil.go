// Package for solving advent of code puzzle 1
package puzzleutil

import (
	"bufio"
	"fmt"
	"os"
)

// brief, Open module mass file and read in all component masses
// param filename, A full or relative path to a mass module file
func ReadMassModuleFile(filename string) int {
	fmt.Println("In ReadMassModuleFile\n")
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error occurred while reading in strings from file: " + filename)
	}

	return 10
}

// brief, Given a value this function calculates expected fuel cost
// param val, A unit to use for calculating fuel cost
func CalculateFuel(val int) int {
	return val
}
