package main

import (
	"fmt"
	//"os"
	"puzzleutil"
	"strconv"
	"stringutil"
)

func main() {
	s := "hello, world\n"
	//file, openError := os.Open("resources/input.txt")
	//var line []byte
	//numRead, readError := os.Read(line)
	fmt.Printf(s)
	fmt.Printf(strconv.Itoa(puzzleutil.CalculateFuel(10)))
	fmt.Printf(stringutil.Reverse(s))
}
