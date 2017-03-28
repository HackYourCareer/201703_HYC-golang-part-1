package main 
import "fmt" 

func main() {
	var words [2]string
	words[0] = "hello"
	words[1] = "world"
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	//cannot use primes (type [6]int) as type [5]int in argument to sum
	//sum(primes)
	// cannot use primes (type [6]int) as type []int in argument to avgl
	//avg(primes)
}

func sum(arr [5]int) int {
	return 123 
}
func avg(arr []int) float64 {
	return 0.0
}
