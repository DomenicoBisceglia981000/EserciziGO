package main
import "fmt"
func main(){
	var n int
	fmt.Println("Inserisci un numero intero")
	fmt.Scan(&n)
	for i := 0; i < n; i++{
		for k := 0; k < i; k++{
			fmt.Print(" ")
		}
		fmt.Print("*\n")
	}
}