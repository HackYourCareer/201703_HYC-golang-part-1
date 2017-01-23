package main

type Player struct {
	FirstName string
	LastName string
}

func (p *Player) Name() {
	println(p.FirstName)
}

func main(){
	p := new(Player)
	p.FirstName = "Piotr"
	p.LastName = "Mścichowski"

	p.Name()
}
