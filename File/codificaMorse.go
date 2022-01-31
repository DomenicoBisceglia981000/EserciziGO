package main
import "fmt"
import "os"
/*

Dopodiché il programma legge da standard input del testo e ne stampa la sua codifica in Morse, stampando una parola per riga,
e separando la codifica morse di un carattere dalla successiva con /.
Il programma non fa distinzione tra maiuscole e minuscole.
Il programma termina con EOF (invio + ctrl D da tastiera)
*/
func main(){
  var s string
  var sFin string
  for{
    fmt.Println("Inserisci testo da convertire in morse : ")
    fmt.Scan(&s)
    file, _ := os.Open("morse.txt")
    b := bufio.NewScanner(file)
    b.Split(bufio.ScanLines)
    for _, c := range s{
      for b.Scan(){
        k:= b.text()
        if c == k[0]{

        }
    }
    }



}
