// Package for solving advent of code puzzle 1
package puzzleutil

import (
	"bufio"
	"fmt"
	//"math"
	"os"
	"strconv"
	"strings"
)

// brief, Open module mass file and read in all component masses
// param filename, A full or relative path to a mass module file
// returns, a byte array of integers
func ReadMassModuleFile(filename string) (numbers []int) {
	fmt.Println("In ReadMassModuleFile\n")
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		read_line := strings.TrimSuffix(scanner.Text(), "\n")
		val, ok := strconv.Atoi(read_line)
		if ok != nil {
			fmt.Fprintf(os.Stderr, "Error occurred while converting string to integer. %s", err)
		}
		//fmt.Printf("val: %d\n", val)
		numbers = append(numbers, val)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr,
			"Error occurred while reading in strings from file: %s", filename)
	}

	return numbers
}

// brief, Given a value this function calculates expected fuel cost
// param val, A unit to use for calculating fuel cost
func CalculateFuel(val int) int {
	return 10 //math.Floor(val/3) - 2
}
