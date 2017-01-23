package main 
import "fmt"

func main() {
	src := []string{"a", "b", "c", "d", "e"}
	fmt.Println("src:", src)

	middle := src[1:4]
	fmt.Println("middle:", middle)

	first := src[:2]
	fmt.Println("first:", first)

	all := src[:]
	fmt.Println("all:", all)

	all[2] = "X"
	fmt.Println("all:", all)
	fmt.Println("src:", src)
}