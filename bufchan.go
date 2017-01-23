package main 

func main() {
	bufferedChan := make(chan string, 2)
	bufferedChan <- "msg1"
	bufferedChan <- "msg2"
	//bufferedChan <- "msg3" // deadlock
}