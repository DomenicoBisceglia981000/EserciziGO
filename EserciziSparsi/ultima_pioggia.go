package main
import "fmt"
func main(){
	var giorno, mm, uGiorno, umm int
	giorno = 1
	for{
		fmt.Print("Inserisci mm di pioggia del giorno ", giorno, " : ")
		fmt.Scan(&mm)
		if mm > 0{
			uGiorno = giorno
			umm = mm
		}
		giorno ++
	}
	fmt.Print("L'ultimo giorno in cui ha piovuto è stato il giorno : ", uGiorno, " con ", umm, " mm di pioggia")
}