package main
import "fmt"
import "unicode"
func main(){
  var s string
  fmt.Print("Inserisci una parola senza spazi: ")
  fmt.Scan(&s)
  var maiu, minu, cons, voc int
  for _, c := range s{
    if èLetteraValida(c){
      if èMaiuscola(c){
        maiu++
      }else{
        minu++
      }
      if èVocale(c){
        voc++
      }else{
        cons++
      }
    }
  }
  fmt.Println("Le lettere maiuscole sono : ", maiu)
  fmt.Println("Le lettere minuscole sono : ", minu)
  fmt.Println("Le consonanti sono : ", cons)
  fmt.Println("Le vocali sono : ", voc)
}
func èLetteraValida(l rune) bool{
  if unicode.IsLetter(l){
    return true
  }else{
    return false
  }
}
func èMaiuscola(l rune) bool{
  if unicode.IsUpper(l){
    return true
  }else{
    return false
  }
}
func èVocale(l rune) bool{
  if l == 65 || l == 97 || l == 69 || l == 101 || l == 73 || l == 105 || l == 79 || l == 111 || l == 85 || l == 117{
    return true
  }else{
    return false
  }
}
