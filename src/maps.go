package main 
import "fmt"

func main() {
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