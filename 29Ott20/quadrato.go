/*

2. Scrivere un programma quadrati.go che legge da standard input
un intero n positivo e calcola il massimo quadrato (k^2) <= n.

Per calcolare il quadrato di un numero n usate n*n.

---

*/
package main
import "fmt"
func main(){
	var n int
	var k = 0
	fmt.Scan(&n)
	for{
		k ++
	
		if k*k > n{
			break
		}

	}
	fmt.Println(k - 1)


}