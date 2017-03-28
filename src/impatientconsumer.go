package main 
import "fmt"
import "sync"
import "time"

func main() {
	msgChan := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go impatientConsumer("Consumer",msgChan, &wg)
	wg.Wait()
}

func impatientConsumer(name string, msgChan chan string, wg *sync.WaitGroup) {
	for {
		select {
		case msg := <-msgChan:
			fmt.Printf("[%s] Received message: %s\n", name, msg)
			if msg == DONE_MSG {//OMIT
				wg.Done()//OMIT
				return//OMIT
			}//OMIT
		case <-time.After(time.Second * 3):
			fmt.Println("I am useless, bye bye!")
			wg.Done()
			return
		}
	}
}