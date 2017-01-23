package main
import "os"
import "fmt"

func main() {
	f, err := os.Open("abc.txt")
	if err != nil {
		fmt.Println("Got error", err)
		return
	}
	defer f.Close() // HL
	f.WriteString("sth") // HL

	i := 10
	defer cleanup(i)
	i = 20
	defer cleanup(i)
}

func cleanup(i int) {
	fmt.Printf("Cleanup %d\n", i)
}