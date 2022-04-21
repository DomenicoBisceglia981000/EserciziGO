/*
scrivere un programma che controlla se un punto (x1,y1) sta sopra, sotto o appartiene ad una retta y = mx + q data, a meno di una differenza pari a 10^-6.
Si dichiarino i dati in ingresso come float64.
stampa "sopra" "sotto" o "sulla" retta
y1 < = > m*x1 + q
*/
package main
import "fmt"
import "math"
func main(){
	var x1, y1, m, q float64
	fmt.Println("Inserisci le coordinate di un punto : ")
	fmt.Scan(&x1, &y1)
	fmt.Println("Inserisci i valori dei coefficenti m e q")
	fmt.Scan(&m, &q)
	if (math.Abs(y1)  < 1e-6) == (math.Abs(m*x1 + q) < 1e-6){
		fmt.Println("appartiene alla retta")
		
	}else if y1 < (math.Abs(m*x1 + q) < 1e-6){
		fmt.Println("sotto la retta")
	}else {
		fmt.Println("sopra la retta")
	}

}

//NON FUNZIONANTE! DA SISTEMARE!!!!!