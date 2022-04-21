package main
import "fmt"
func main(){
	var n, d int
	
	fmt.Print("Inserisci un intero : ")
	fmt.Scan(&n)
	d = n
	for i := 0; i < n; i++{
		for k := 0; k < i; k++{
			fmt.Print(" ")
		}
		fmt.Print("*")
		for y := 0; y <= d; y ++{
			fmt.Print(" ")
		}
		if i < n -1{
			fmt.Print("*\n")
		}

		d = d - 2


	}
}