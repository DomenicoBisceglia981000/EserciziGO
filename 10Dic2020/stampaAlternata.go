package main
import (
  "bufio"
  "os"
  "fmt"
)

type Position struct{
  Filename string 
    Offset int    
    Line int    
    Column int
}

func main(){
  var file Position
  file.Line = 1
  fileName:= os.Args[1]
  f, err:= os.Open(fileName)
   if err != nil{
    fmt.Printf("Errore nell'apertura del file: %v\n", err)
    return
  }
file.Filename = fileName
  defer f.Close()
b:= bufio.NewScanner(f)
c:= bufio.NewScanner(f)
for b.Scan(){
  if file.Line %2 == 0{
    line:= b.Text()
  fmt.Printf("Line pari: <%s>\n", line)
  }
  file.Line++
}
file.Line = 1
for c.Scan(){
  if file.Line %2 != 0{
    linedis:= c.Text()
  fmt.Printf("Line dispari: <%s>\n", linedis)
  }
  file.Line++
}

}
