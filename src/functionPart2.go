package main

import "fmt"

//START OMIT
func namedValues(val int) (x, y int) {
	x = val * 2
	y = val * 3
	return x,y
}

func noRetArg(val int) (x, y int) {
	x = val * 2
	y = val * 3
	return
}

func main() {
	fmt.Println(namedValues(2))
	fmt.Println(noRetArg(2))
}
//END OMIT
