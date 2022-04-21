/*
scrivere un programma che verifica se due frazioni num1/den1 e num2/den2 sono equivalenti e stampa "equivalenti" o "non equivalenti", a seconda del caso. Usare il tipo int per i dati in input. 
Pensare a una soluzione che non facci auso di float64
*/
package main
import "fmt"
func main(){
	var num1, num2, den1, den2 int
	fmt.Print("Inserisci num e den fraz1 : ")
	fmt.Scan(&num1, &den1)
	fmt.Print("Inserisci num e den fraz2 : ")
	fmt.Scan(&num2, &den2)
	if (num1 * den2) == (num2 * den1){
		fmt.Println("frazioni equivalenti")
	}else{
		fmt.Println("frazioni non equivalenti")
	}
}