/*
Scrivere un programma lunghezza_tot.go che legge da standard input un int totLen e una sequenza di stringhe (una per riga) e ne somma le 
lunghezze fino a raggiungere (o superare) totLen. 
Quanto termina, stampa la somma delle lunghezze delle stringhe lette.

*/
package main
import "fmt"
func main(){
	var totLen int
	var s string
	fmt.Println("Inserisci un valore per la lunghezza totale : ")
	fmt.Scan(&totLen)
	var i int
	for i < totLen{
		fmt.Print("Inserisci una parola")
		fmt.Scan(&s)
		i = i + len(s)
	}
	fmt.Println(i)
}