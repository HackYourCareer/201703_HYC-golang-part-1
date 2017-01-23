package main 
import "fmt" 
import "sync"
import "time"

const DONE_MSG = "DONE!"

func main() {
	msgChan := make(chan string)
	wg := sync.WaitGroup{}
	go producer("Producer",msgChan,&wg)
	wg.Add(1)
	go consumer(fmt.Sprint("Consumer"), msgChan)
	wg.Wait()
	msgChan <- DONE_MSG
}

func producer(name string, msgChan chan string, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		msgChan <- fmt.Sprintf("[%s] Message:%d", name, i+1)
		time.Sleep(time.Second)
	}
	fmt.Printf("[%s] Done\n", name)
	wg.Done()
}

func consumer(name string, msgChan chan string) {
	got := 0
	for {
		msg := <-msgChan
		got++
		fmt.Printf("[%s] Received message nr %d: %s\n", name, got, msg)
		if msg == DONE_MSG {
			return
		}
	}
}