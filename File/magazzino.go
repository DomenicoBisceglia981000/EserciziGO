package main
import "fmt"
import "os"

type Prodotto struct{
  codice, descr string
  prezzo float64
}

func main(){
  var prodotto Prodotto
  listino := make(map[string]Prodotto)
  nomefile := os.Args[1]
  file, _ := os.Open(nomefile)
  b := bufio.NewScanner(file)
  b.Split(bufio.ScanLines)
  for b.Scan(){
    s:= b.text()
    if string(s[0]) != "+" && string(s[0]) != "-"{
      prodotto.codice = strings.Split(s, ",")[0]
      prodotto.descr = strings.Split(s, ",")[1]
      prezzo, _ := strconv.parseFloat(strings.Split(s, ",")[2]),

    }
  }
}
