/*
Scrivere un programma cesare.go che legge da standard input 
un valore intero non negativo k (la chiave di cifratura) e
una sequenza di lettere minuscole consecutive 
(sulla stessa riga e senza spazi) terminate da <invio> ('\n').
Il programma stampa la sequenza letta cifrata secondo il 
cifrario di Cesare, usando come chiave k (quella fornita dall'utente):
ogni lettera del testo in chiaro è sostituita nel 
testo cifrato dalla lettera che si trova k posizioni dopo 
nell'alfabeto, ritornando dopo la zeta alla lettera a.

*/
package main
import (
	"fmt"
	"unicode"
	)	
func main(){
	var k rune
	var s string
	var sok bool
	sok = false
	for k <= 0{
		fmt.Print("Inserisci un numero : ")
		fmt.Scan(&k)
	}
	for sok == false{
		fmt.Print("Inserisci la parola da cifrare : ")
		fmt.Scan(&s)
		for _, c := range s{
			if unicode.IsUpper(c){
				sok = false
				break
			}else{
				sok = true
			}
		}
	}
	
	for _, c := range s{
		c = c + k
		for c > 'z'{
			c = c - 26		
		}

		fmt.Print(string(c))


	}
}