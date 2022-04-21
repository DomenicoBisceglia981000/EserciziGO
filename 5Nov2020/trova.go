/*
SPECIFICHE: Scrivere un programma trova.go che legge da standard input un carattere (rune)
 e una stringa e stampa la posizione del carattere nella stringa, o -1 se il carattere non c'è.
*/
package main
import "fmt"
func main(){
	var ch rune
	var s string
	var pos int
	fmt.Println("Inserisci un carattere : ")
	fmt.Scanf("%c", &ch)
	fmt.Println("Inserisci una stringa : ")
	fmt.Scan(&s)
	for i, c:= range s {
		if ch == c{
			pos = i + 1
		}
	}
	if pos == 0{
		fmt.Println(-1)
	}else{
		fmt.Println(pos)
	}

	
}