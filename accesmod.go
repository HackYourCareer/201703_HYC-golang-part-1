package mylib
import "fmt"
func Sum(x, y int) int {
	return x + y
}

func save(val string) bool {
	...
}


package main
import (
"myLib"
"fmt"
)

func main() {
	fmt.Println(mylib.Sum(2, 3))
	mylib.save()
}


