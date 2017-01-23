package main

import (
	"strconv"
	"errors"
	"fmt"
)
//START OMIT
func assignMe(key string)(int, error){
	intVal, err := strconv.Atoi(key)

	if err != nil {
		return 0,errors.New("A coś Ty mi tu nawpuszczał Wąski ?! ")
	}

	return intVal*len(key), nil
}

func main() {
	funcVal := assignMe

	value, err := funcVal("młode wilki")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(value)
}
//END OMIT
