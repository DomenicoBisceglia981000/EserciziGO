/*
scrivere un programma che legge un intero n e a seconda del suo valore stampa uno dei messaggi "n è pari" oppure "n è dispari"
*/
package main
import "fmt"
func main(){
	var n int
	fmt.Print("Inserisci un numero : ")
	fmt.Scan(&n)
	if (n % 2) == 0{
		fmt.Println("n è pari")
	}else{
		fmt.Println("n è dispari")
	}
}