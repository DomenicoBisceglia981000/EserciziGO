/*
 Scrivere un programma max_char.go che legge da standard input una sequenza di 5 caratteri ASCII (byte) e stampa il maggiore in ordine lessicografico (codice ASCII).

*/
package main
import "fmt"
func main(){
	var ch byte
	var max byte
	fmt.Println("Inserisci una sequenza di 5 caratteri")
	for i := 0; i < 5; i++{
		fmt.Scanf("%c", &ch)
		if ch > max{
			max = ch
		}
	}
	fmt.Println(string(max))
}