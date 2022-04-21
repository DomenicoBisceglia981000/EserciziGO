/*
scrivere un programma che legga due numeri interi e ne stampi il maggiore
*/
package main
import "fmt"
func main(){
	var min, max int
	fmt.Print("Inserisci due numeri : ")
	fmt.Scan(&min, &max)
	if min > max {
		min, max = max, min
	}
	fmt.Println("Il numero maggiore è :", max)
}