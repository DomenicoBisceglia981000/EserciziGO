/*
Scrivere un programma che legge un numero intero. se il numero non è compreso tra 0 e 30, stampa "voto non valido", altrimenti non stampa niente
*/
package main
import "fmt"

func main(){
	var voto int
	fmt.Print("Inserisci un voto : ")
	fmt.Scan(&voto)

	if((voto <= 0) || (voto >= 30)){
		fmt.Println("voto non valido")
	}
}