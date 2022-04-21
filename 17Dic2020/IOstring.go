package main
import ("fmt"
		"os")
func main(){
	var gg int
	var mm int
	var  aaaa int
	data := os.Args[1]
	fmt.Sscanf(data, "%d-%d-%d", &gg, &mm, &aaaa);
	fmt.Printf("Giorno : %d\nMese : %d \nAnno : %d \n", gg, mm, aaaa)
}