package main
import "fmt"
func main(){
  var s string
  fmt.Print("Inserisci una stringa senza spazi : ")
  fmt.Scan(&s)
  var sm string
  for i, c := range s{
    if i < len(s){
      sm = sm + string(c) + " "
    }else{
      sm = sm + string(c)
    }

  }
  fmt.Println("Stringa spaziata : ", sm)
}
