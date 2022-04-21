/*
1. Scrivere un programma euclide.go che legge da standard input due interi 
a e b, con a >= b, e calcola il MCD tra i due numeri con l'algoritmo di Euclide.

Algoritmo di Euclide:

Dati due numeri naturali a e b, 
1. si controlla se b e` zero. 
2. se lo e`, a è il MCD. 
3. se non lo è, si assegna ad r il resto della divisione a / b, 
   si pone a = b e b = r e si ripete da 1.



*/
package main
import "fmt"
func main(){
	var a, b, r int
	fmt.Scan(&a)
	fmt.Scan(&b)
	if a < b{
		c := a
		a = b
		b = c
	}
	for{
		if b == 0{
			fmt.Println("MCD è : ", a)
			break
		}
		r = a % b
		a = b
		b = r
	}

}