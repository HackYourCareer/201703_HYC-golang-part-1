package main

import "fmt"

func main(){
	func (values ... int) {
		fmt.Println("I've been called with : ", values)
	}()
}
