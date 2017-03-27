package main
import "fmt"
///START OMIT
type Person struct {
	name string
	lastName string
}

//ptr is a method receiver
func (ptr *Person) modifyByPtr() {
	ptr.name = "Tom"
	ptr.lastName = "Hanks"
}
func (p Person) modifyByValue() {
	p.lastName = "Wick"
	p.name = "John"
}
func main() {
	me := Person{
		name: "Piotr",
		lastName: "Mścichowski",
	}
	fmt.Println(me)
	ptr := &me
	ptr.modifyByPtr()
	fmt.Println(me)
	me.modifyByValue()
	fmt.Println(me)
}
//END OMIT