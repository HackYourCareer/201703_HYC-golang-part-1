package main 
import "fmt"
import "time"

func main() {
	go longComputation(1000)
	fmt.Println("Will be printed immediately")
}

func longComputation(x int) {
	time.Sleep(time.Hour)
}