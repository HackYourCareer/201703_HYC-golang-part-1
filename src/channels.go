package main 

import (
	"fmt" 
	// "time"
)

func main() {
	msgChan := make(chan string)

	go func() {
		msgChan <- "message"
	}()
	go func() {
		val := <-msgChan
		fmt.Println(val)
	}()
	 // time.Sleep(time.Second)	
	 // msgChan <- "message 2"  

}