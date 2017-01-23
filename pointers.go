package main
import "fmt"
///START OMIT
type Person struct {
	name string
	lastName string
}
func modify(ptr *Person) {
	ptr.name = "Tom"
	ptr.lastName = "Hanks"
}
func mod(p Person){
	p.lastName = "`Wick"
	p.name = "John"
}
func main() {
	me := Person{
		name: "Piotr",
		lastName: "Mścichowski",
	}
	fmt.Println(me)
	modify(&me)
	fmt.Println(me)
	mod(me)
	fmt.Println(me)
}
//END OMIT