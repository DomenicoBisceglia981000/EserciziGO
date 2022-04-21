package main
import "fmt"
func main(){
const k = 5
	var n int
	for i := 0; i < k; i ++{
		fmt.Println("inserisci numero : ")
		fmt.Scan(&n)
		fmt.Println(n * 2)
	}
}