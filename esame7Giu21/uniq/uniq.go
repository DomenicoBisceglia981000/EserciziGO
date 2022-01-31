package main
import(
  "fmt"
  "os"
  "bufio"
)

func main(){
  if len(os.Args) == 2{
    filename := os.Args[1]
    file, err := os.Open(filename)
    if err == nil{
      scanner := bufio.NewScanner(file)
      defer file.Close()
      scanner.Split(bufio.ScanLines)
      precedente := " "
      conteggioRighe := 1
      for scanner.Scan(){
        riga := scanner.Text()
        if precedente == " "{
          precedente = riga
          continue
        }else{
          if riga != precedente {
            fmt.Println(conteggioRighe, precedente)
            conteggioRighe = 1
          }else{
            conteggioRighe++
          }
          precedente = riga
        }
      }
      //STAMPO ANCHE L'ULTIMA RIGA CHE ALTRIMENTI VERREBBE OMESSA, MA SOLO SE E' SINGOLA, ALTRIMENTI NON AVREI PROBLEMI
      fmt.Println(conteggioRighe, precedente)
    }

  }
}
