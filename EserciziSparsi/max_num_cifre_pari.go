package main
import "fmt"
func main(){
	var s string
	var cont, contMax int
	for{
		fmt.Print("Inserisci una sequenza di numeri : ")
		fmt.Scan(&s)
		if s == "x"{
			break
		}
		for _, c := range s{
			if int(c) % 2 == 0{
				cont ++
			}
		}
		if cont > contMax{
			contMax = cont
		}
		cont = 0
	}
	fmt.Print("Il numero massimo di numeri pari all'interno di una delle sequenze è : ", contMax)
}