package main
import "fmt"

func main() {
	// insert in the middle
	src := []string{"a", "b", "d", "e"}
	fmt.Println(src)
	src = append(src[:2], append([]string{"c"}, src[2:]...)...)
	fmt.Println("src:", src)

	//delete
	copy(src[2:], src[3:])
	fmt.Println("src after copy:", src)
	src[len(src) - 1] = ""
	src = src[:len(src) - 1]
	fmt.Println("src:", src)
}