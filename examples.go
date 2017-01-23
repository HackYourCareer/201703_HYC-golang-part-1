package main

import (
	"fmt"
	"sync"
	"time"
	"strconv"
	"os"
)

func main() {
	slideSlices()
}

// -------------- ARRAYS -----------------
func slideArray() {
	var arr [2]string
	arr[0] = "hello"
	arr[1] = "world"
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	//cannot use primes (type [6]int) as type []int in argument to sum
	//sum(primes)
	// cannot use primes (type [6]int) as type []int in argument to avgl
	//avg(primes)
}

func sum(arr []int) int {
	// TODO
	return 123
}

func avg(arr []int) float64 {
	//TODO
	return 0.0
}

// -------------- SLICES -----------------
func slideSlices() {
	var sl []string // HL
	fmt.Printf("sl len: %d, cap: %d\n", len(sl), cap(sl))

	slM := make([]string, 5) // HL
	fmt.Printf("slM len: %d, cap: %d\n", len(slM), cap(slM))

	// sl[0] = "a" // panic: runtime error: index out of range
	slM[0] = "a" // OK
	// slM[10] = "b" // panic: runtime error: index out of range

	sl = append(sl, "first") // HL
	fmt.Println(sl)

	slM = append(slM, "last")
	fmt.Printf("slM after append: %v, len: %d, cap: %d\n", slM, len(slM), cap(slM))

}

func slideSlicing() {
	src := []string{"a", "b", "c", "d", "e"}
	fmt.Println("src:", src)

	middle := src[1:4]
	fmt.Println("middle:", middle)

	first := src[:2]
	fmt.Println("first:", first)

	all := src[:]
	fmt.Println("all:", all)

	all[2] = "X"
	fmt.Println("all:", all)
	fmt.Println("src:", src)
}

func slideSliceTrickyPart() {
	// insert in the middle
	src := []string{"a", "b", "d", "e"}
	fmt.Println(src)
	src = append(src[:2], append([]string{"c"}, src[2:]...)...)
	fmt.Println("src:", src)

	//delete
	copy(src[2:], src[3:])
	fmt.Println("src after copy:", src)
	src[len(src) - 1] = ""
	src = src[:len(src) - 1]
	fmt.Println("src:", src)
}

// -------------- maps -----------------

func slideMaps() {
	worldChampions := map[int]string{
		1994:"Brasil",
		1998:"France",
		2002:"Brasil",
		2006:"Italy",
		2010:"Spain",
	}
	worldChampions[2014] = "Germany"
	fmt.Println(worldChampions[1994])
	_, ex := worldChampions[1997]
	fmt.Println(ex)
	delete(worldChampions, 2006)
	fmt.Println(worldChampions)

}

// -------------- for, if, defer -----------------

func slideFor() {
	for i := 0; i < 10; i++ {

	}

	flag := true
	for {
		if flag {
			break
		}
	}

	str := "12"

	if n, err := strconv.Atoi(str); err != nil {
		fmt.Println("Got error when parsing ", str)
	} else {
		fmt.Printf("Number is %d\n", n)
	}
	// n = 14 - Unresolved reference
}

func slideDefer() {
	f, err := os.Open("abc.txt")
	if err != nil {
		fmt.Println("Got error", err)
		return
	}
	defer f.Close()

	f.WriteString("sth")

	i := 10
	defer cleanup(i)
	i = 20
	defer cleanup(i)
}

func cleanup(i int) {
	fmt.Printf("Cleanup %d\n", i)
}

func slideSwitch() {
	var todayMeme string
	switch today := time.Now().Weekday(); today {
	case time.Monday:
		fallthrough
	case time.Tuesday:
		todayMeme = "grumpyCat"
	case time.Friday:
		todayMeme = "weekend"
	default:
		todayMeme = "randomMeme"
	}
	fmt.Println(todayMeme)

}

// -------------- GORUTINES, CHANNELS -----------------

const DONE_MSG = "DONE"

func caseNoProducer() {
	msgChan := make(chan string)
	go consumer(fmt.Sprint("Consumer"), msgChan)

}

func caseProdConsumer() {
	msgChan := make(chan string)
	wg := sync.WaitGroup{}

	go noProducer(&wg)
	wg.Add(1)
	//for i := 0; i < 2; i++ {
	//	go slowProducer(fmt.Sprintf("Producer  %d", i), msgChan, &wg)
	//	wg.Add(1)
	//}
	//go iDontCareProducer("IdontCare",msgChan,&wg)
	//wg.Add(1)
	go consumer(fmt.Sprint("Consumer"), msgChan)
	wg.Wait()
	//msgChan <- "DONE"
}

func caseProdConsumerWithBufferedChan() {

}

func noProducer(wg *sync.WaitGroup) {
	time.Sleep(time.Second * 15)
	wg.Done()
}

func slowProducer(name string, msgChan chan string, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		msgChan <- fmt.Sprintf("[%s] Message:%d", name, i+1)
		time.Sleep(time.Second)
	}
	fmt.Printf("[%s] Done\n", name)
	wg.Done()
}

func iDontCareProducer(name string, msgChan chan string, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		select {
		case msgChan <- fmt.Sprintf("[%s] Not so important message:%d", name, i+1):
		default:
			fmt.Println("I don't care...")
		}
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

func impatientConsumer(name string, msgChan chan string) {
	got := 0
	for {
		select {
		case msg := <-msgChan:
			got++
			fmt.Printf("[%s] Received message nr %d: %s\n", name, got, msg)
			if msg == DONE_MSG {
				return
			}
		case <-time.After(time.Second * 5):
			fmt.Println("I am useless, bye bye!")
			return
		}

	}
}
