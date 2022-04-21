/*
Convertitore text2Morse
-----------------------

Scrivere un programma morse.go per la traduzione in alfabeto Morse di testo. 
Il programma legge dal file morse.txt la codifica Morse dei caratteri
nel seguente formato:

A . - 
B - . . . 
C - . - . 

Dopodiché il programma legge da standard input del testo e ne stampa la sua codifica in Morse, stampando una parola per riga, e separando la codifica morse di un carattere dalla successiva con /. 
Il programma non fa distinzione tra maiuscole e minuscole.
Il programma termina con EOF (invio + ctrl D da tastiera)

Esempio
-------

$ go run morse.go
ciao
- . - .  / . .  / . -  / - - -  / 
ctrl D
*/
package main
import ("fmt"
		"os"
		"bufio")
func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var parola string
	var morse [] string
	fmt.Print("Inserisci una parola da tradurre in morse : ")
	fmt.Scan(&parola)
	for _, c := range parola{
		for scanner.Scan(){
			
		}
	}
}
