package main
import "fmt"
import "unicode"

func main(){
  var s, s2 string
  fmt.Println("Inserisci una parola senza spazi : ")
  fmt.Scan(&s)
  //toMinu
  for _,c := range s{
    if unicode.IsUpper(c){
      s2 = s2 + string(inMinuscolo(c))
    }else{
      s2 = s2 + string(c)
    }
  }
  fmt.Println("In minuscolo : ", s2)
  s2 = ""
  //ToMaiu
  for _,c := range s{
    if unicode.IsLower(c){
      s2 = s2 + string(inMaiuscolo(c))
    }else{
      s2 = s2 + string(c)
    }
  }
  fmt.Println("In maiuscolo : ", s2)


}
func inMaiuscolo(car rune) rune{
  car = unicode.ToUpper(car)
  return car
}
func inMinuscolo(car rune) rune{
  car = unicode.ToLower(car)
  return car
}
