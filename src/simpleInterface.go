package main

import "fmt"

type Player interface {
	Hello(message string)
}

type Test struct {
}

func (t Test) Hello(message string) {
	fmt.Println("Hello", message)
}
