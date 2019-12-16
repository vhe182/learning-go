package main

import (
	"fmt"
	"stringutil"
)

func main() {
	s := "hello, world\n"
	fmt.Printf(s)
	fmt.Printf(stringutil.Reverse(s))
}
