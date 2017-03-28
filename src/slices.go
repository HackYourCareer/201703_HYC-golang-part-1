package main
import "fmt"

func main() {
	var emptySlice []string
	fmt.Printf("empty slice len: %d, cap: %d\n", len(emptySlice), cap(emptySlice))

	letters := make([]string, 5, 10)
	fmt.Printf("letters len: %d, cap: %d\n", len(letters), cap(letters))

	// emptySlice[0] = "a" // panic: runtime error: index out of range
	letters[0] = "a" // OK
	// letters[10] = "b" // panic: runtime error: index out of range

	letters = append(letters, "f")
	fmt.Printf("letters after append: %v, len: %d, cap: %d\n", letters, len(letters), cap(letters))
}