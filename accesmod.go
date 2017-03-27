Name := "Piotr" //Exported
id := "1kh2uss73x" //Unexported

//function sayHello is unexported
func sayHello(){ 
	fmt.Println("Hello") //fmt.Println is exported
}

//Exported
func ExporteMe(){

}
