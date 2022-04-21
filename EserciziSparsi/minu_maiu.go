package main
import (
	"fmt"
	"unicode"
	)
func main(){
	var s string
	minu, maiu := false, false
	fmt.Print("Inserisci una parola : ")
	fmt.Scan(&s)
	for i, c := range s{
		if unicode.IsLower(c) {
			minu = true
		}else{
			maiu = true
		}
		_ = i
	}
	if minu && maiu{
		fmt.Print("Sia minuscole che maiuscole")
	}else if minu{
		fmt.Print("Solo minuscole")
	}else{
		fmt.Print("solo Maiuscole")
	}
}