package main
import "fmt"
func main(){
	const k = 5
	var n, ris int
	for i := 0; i < k; i ++{
		fmt.Println("Inserisci numero : ")
		fmt.Scan(&n)
		ris = ris + n
	}
}