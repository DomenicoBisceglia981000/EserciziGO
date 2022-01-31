package main
import(
  "fmt"
  "os"
  "bufio"
  "strings"
  "sort"
)

func printKeywords(listaKey map[string][]int){
  chiavi := make([]string, 0)
  for i := range listaKey{
    chiavi = append(chiavi, i)
  }
  sort.Strings(chiavi)
  for i := 0; i < len(chiavi); i++{
    fmt.Print(chiavi[i], ": ", listaKey[chiavi[i]], "\n")
  }
}

func main(){
  if len(os.Args) == 2{
    filename := os.Args[1]
    file, err := os.Open(filename)
    if err == nil{
      //creo il primo scanner che leggerà riga per riga il mio file
      scanner := bufio.NewScanner(file)
      defer file.Close()
      scanner.Split(bufio.ScanLines)
      mappaKeywords := make(map[string][]int)
      contaRighe := 0
      for scanner.Scan(){
        contaRighe++
        riga := scanner.Text()
        b := bufio.NewScanner(strings.NewReader(riga))
        //creo un secondo scanner che analizza ogni parola della riga letta da scanner.Text()
        b.Split(bufio.ScanWords)
        for b.Scan(){
          parola := b.Text()
          switch parola{
            case "break", "case", "chan", "const", "continue", "default", "defer", "else", "fallthrough", "for", "func", "go", "goto", "if", "import", "interface", "map", "package", "range", "return", "select", "struct", "switch", "type", "var" :
              mappaKeywords[parola] = append(mappaKeywords[parola], contaRighe)
            default :
              continue
          }
        }
      }
      printKeywords(mappaKeywords)
    }else{
      fmt.Println("File not found")
    }


  }else{
    fmt.Println("A file name, please")
  }
}
