package main

import "fmt"
import "strconv"

func main() {
	str := "12"
	if n, err := strconv.Atoi(str); err != nil {
		fmt.Println("Got error when parsing ", str)
	} else {
		fmt.Printf("Number is %d\n", n)
	}
	// n = 14 - Unresolved reference
}