package main

import "fmt"

func main() {
	//PARTE 1
	var ch rune
	fmt.Print("Inserisci un carattere : ")
	fmt.Scanf("%c", &ch)
	fmt.Println(string(ch), "è un carattere.")
	fmt.Println(string(ch-1), ";", string(ch), ";", string(ch+1))

	if (ch >= 'A' && ch <= 'L') || (ch >= 'a' && ch <= 'l') {
		fmt.Println(string(ch), "[a,l]")
	} else {
		fmt.Println(string(ch), "non è compreso tra A e L")
	}

	//PARTE 2
	var s string
	fmt.Print("Inserisci una stringa :")
	fmt.Scan(&s)
	for _, c := range s {
		fmt.Println(string(c))
	}
}
