/*
Scrivere un programma trend.go che legge da standard input una sequenza di 7 interi e 
stampa "+" quando il nuovo numero letto è maggiore del precedente, stampa "-" quando è minore e stampa "=" quando è uguale.
*/
package main
import "fmt"
func main(){
	var n int
	fmt.Println("Inserisci un numero : ")
	fmt.Scan(&n)
	var nu int
	for i := 0; i < 6; i++{
		fmt.Println("Inserisci un numero : ")
		fmt.Scan(&nu)
		if nu > n{
			fmt.Println(">")
		}else if nu == n{
			fmt.Println("=")
		}else if nu < n{
			fmt.Println("<")
		}
		n = nu
	} 
}