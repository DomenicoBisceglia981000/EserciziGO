package main
import "fmt"
func main(){
	var n int
	tot, som := 1, 0
	fmt.Print("Inserisci intero : ")
	fmt.Scan(&n)
	for i := 0; i <n; i ++{
		
		for z := 0; z < tot; z++{
			fmt.Print("*")
		}
		fmt.Println()
		som = tot
		tot = tot + som
	}
}