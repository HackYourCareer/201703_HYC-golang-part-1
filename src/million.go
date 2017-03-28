package main

import "sync"
import "time"
import "fmt"
import "math/rand"

func main() {
	wg := sync.WaitGroup{}
	bef := time.Now()
	const len = 1000000 // million
	wg.Add(len)
	var out [len]int32
	for i := 0; i < len; i++ {
		go func(idx int) {
			out[idx] = rand.Int31n(1000)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(bef))
}
