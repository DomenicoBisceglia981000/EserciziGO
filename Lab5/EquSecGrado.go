package main
import "fmt"
import "math"
import "math/cmplx"
import "os"
import "strconv"
func main(){
  a, _ := strconv.Atoi(os.Args[1])
  b, _ := strconv.Atoi(os.Args[2])
  c, _ := strconv.Atoi(os.Args[3])
  CalcolaEq(a, b, c)

}
func CalcolaEq(a, b, c int){
  //Calcolo il delta
  delta := (b*b) - (4*a*c)
  switch{
    case delta > 0 :
    //ci sono due soluzioni reali
    ris1 := (-float64(b) + math.Sqrt(float64(delta))) / 2.0*float64(a)
    ris2 := (-float64(b) - math.Sqrt(float64(delta))) / 2.0*float64(a)
    fmt.Println("Due soluzioni reali :", ris1, ris2)
  case delta == 0 :
    //c'è una sola soluzione reale
    ris1 := (-float64(b) + math.Sqrt(float64(delta))) / 2.0*float64(a)
    ris2 := (-float64(b) - math.Sqrt(float64(delta))) / 2.0*float64(a)
    fmt.Println("Due soluzioni coincidenti :", ris1, ris2)
    case delta < 0 :
      //non ci sono soluzioni reali
      ris1 := (complex(float64(-b), 0) + cmplx.Sqrt(complex(float64(delta), 0)))/complex(2*float64(a), 0)
      ris2 := (complex(float64(-b), 0) - cmplx.Sqrt(complex(float64(delta), 0)))/complex(2*float64(a), 0)
      fmt.Println("Due soluzioni immaginarie : ", ris1, ris2)
  }
}
