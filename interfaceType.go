package main

import (
	"fmt"
)
//START OMIT
type test struct {
	val string
}

// interface{} has no methods
func interfaceType(d interface{}){
	fmt.Println(d)
	//work on copy of an interface type
}

func main() {
	interfaceType(4)
	interfaceType(test{"test value"})
}
//END OMIT
