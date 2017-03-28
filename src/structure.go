package main

import "fmt"

type Player struct {
	FirstName string
	LastName string
}

func main(){
	goalKeeper := Player{
		FirstName: "Piotr",
		LastName: "Mścichowski", //-> comma required
	}

	proPlayer := Player{"Cristiano", "Ronaldo"}

	fmt.Println(goalKeeper, proPlayer)
}
