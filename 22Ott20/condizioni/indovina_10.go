/*
scriver eun programma che fissa un numero intero tra 1 e 10 da indovinare, legge un intero da standard input e 
- se il numero in input è fuori dall'intervallo 1-10, stampa "valore non valido"
- se il numero è quello fissato, stampa "hai indovinato"
- altrimenti stampa "non hai indovinato"
*/

package main
import "fmt"

func main(){
	const n = 9
	var numUtente int
	fmt.Print("Indovina il numero compreso tra 1 e 10 : ")
	fmt.Scan(&numUtente)

	if numUtente >= 1 && numUtente <= 10 {
		if numUtente == n{
			fmt.Println("Hai indovinato !!!")
		}else{
			fmt.Println("Non hai indovinato")
		}
	}else{
		fmt.Println("Valore non valido")
	}
}