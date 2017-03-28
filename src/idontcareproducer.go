package main
import "sync"
import "fmt"
import "time"
func main() {
	msgChan := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(2) 
	go func () {
		fmt.Println("Got: ",<-msgChan)
		wg.Done()
	}()
	go iDontCareProducer("Producer",msgChan,&wg)
	wg.Wait()
}
func iDontCareProducer(name string, msgChan chan string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		select { // HL
			case msgChan <- fmt.Sprintf("[%s] Not so important message:%d", name, i+1): // HL 
			case <-time.After(time.Second): fmt.Println("Timeout") // HL
		} 
	}
	fmt.Printf("[%s] Done\n", name)
	wg.Done()
}