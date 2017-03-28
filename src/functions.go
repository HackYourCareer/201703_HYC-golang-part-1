package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func hello(name string) {
	fmt.Println("hello", name)
}
func main() {
	hello("Maciej")
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
