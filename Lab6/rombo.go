package main
import "fmt"
func main(){
  var n int
  fmt.Print("Inserisci un numero : ")
  fmt.Scan(&n)
  diag := 2 * n + 1
  stringaSpazi := ""
  stringaAsterischi := ""
  for i := 0; i < diag; i++{
    stringaSpazi += " "
    stringaAsterischi += "*"
  }
  for i := 0; i < diag/2; i++{
    fmt.Print(stringaSpazi[:diag/2 -i], stringaAsterischi[:2*i +1], "\n")
  }
  fmt.Println(stringaAsterischi)
  for i := diag/2 - 1; i >= 0; i--{
    fmt.Print(stringaSpazi[:diag/2 -i], stringaAsterischi[:2*i+1], "\n")
  }
}
