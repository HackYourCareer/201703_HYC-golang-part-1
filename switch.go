package main 

import "time"
import "fmt"

func main() {
	var todayMeme string
	switch today := time.Now().Weekday(); today { // HL
	case time.Monday:
		fallthrough // HL
	case time.Tuesday:
		todayMeme = "grumpyCat"
	case time.Friday:
		todayMeme = "weekend"
	default: // HL
		todayMeme = "randomMeme"
	}
	fmt.Println(todayMeme)
}