/*
Scrivere un programma conta_vocali.go che legge da standard input una parola e stampa il numero di lettere (rune) che sono vocali (aeiou).
*/
package main
import "fmt"
func main(){
	var s string
	var cont int
	fmt.Print("Inserisci una parola : ")
	fmt.Scan(&s)
	for i, ch := range s{
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'{
			cont ++
		}
		_ = i
	}
	fmt.Println("Ci sono", cont, " vocali nella parola", s)

}