package main
import(
  "fmt"
  "bufio"
  "os"
  "strconv"
  "math"
)

func main(){
  min := math.MaxInt32
  max := math.MinInt32
  somma := 0
  cont := 0
  var media float64
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan(){
    n, err := strconv.Atoi(scanner.Text())
    if err == nil{
      cont++
      somma = somma + n
      if n > max{
        max = n
      }
      if n < min{
        min = n
      }
    }
  }
  media = float64(somma)/float64(cont)
  fmt.Println("min:", min)
  fmt.Println("max:", max)
  fmt.Println("media:", media)
}
