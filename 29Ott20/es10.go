package main
import "fmt"
func main(){
	var h, m int
	for{
		fmt.Println("Inserisci ore e minuti :")
		fmt.Scan(&h, &m)
		if (h >= 0 && h <= 23) && (m >= 0 && m <= 59){
			break
		}
	}
	fmt.Println("ORARIO VALIDO :", h, "/", m)
}