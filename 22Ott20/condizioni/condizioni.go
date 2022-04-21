/*
scrivere un unico programma condizioni.go che testa una ad una le seguenti condizioni :
legge un valore e stampa true o false a seconda che la condizione sia verificata o no
- int uguale a 10
- int diverso da 10
- int diverso da 10 e da 20
- int diverso da 10 o da 20
- int maggiore o uguale a 10
- int compreso tra 10 e venti, nell'intervallo [10,20]
- int compreso tra 10 e venti, nell'intervallo (10,20)
- int compreso tra 10 e venti, nell'intervallo [10,20)
*/
package main
import "fmt"
import "math"

func main(){
    //var a int
    
    /*
    fmt.Print("int uguale a 10: ")
    fmt.Scan(&a)
    fmt.Println(a == 10)
    */

    /*
    fmt.Print("int diverso da 10: ")
    fmt.Scan(&a)
    fmt.Println(a != 10)
    */
    
    /*
    fmt.Print("int diverso da 10 e da 20 : ")
    fmt.Scan(&a)
    fmt.Println((a != 10) && (a != 20))
    */

    /*
    fmt.Print("int diverso da 10 o da 20 : ")
    fmt.Scan(&a)
    fmt.Println((a != 10) || (a != 20))
    */

    /*
    fmt.Print("int compreso tra 10 e 20 nell'intervallo [10, 20] : ")
    fmt.Scan(&a)
    fmt.Println( (a >= 10) && (a <= 20))
    */

    /*
    fmt.Print("int compreso tra 10 e 20 nell'intervallo (10, 20) : ")
    fmt.Scan(&a)
    fmt.Println((a > 10) && (a < 20))
    */

    /*
    fmt.Print("int compreso tra 10 e 20 nell'intervallo [10, 20) : ")
    fmt.Scan(&a)
    fmt.Println( (a >= 10) && (a < 20) )
    */

    /*
    fmt.Print("int esterno all'intervallo [10, 20] : ")
    fmt.Scan(&a)
    fmt.Println((a < 10) || (a > 20))
    */

    /*
    fmt.Print("int tra 10 e 20 (nell'intervallo [10, 20]) o tra 30 e 40 ([30, 40]) : ")
    fmt.Scan(&a)
    fmt.Println( ( (a >= 10) && (a <= 20) ) || ( (a >= 30) && ( a <= 40) ) )
    */

    /*
    fmt.Print("int multiplo di 4 ma non di 100 : ")
    fmt.Scan(&a)
    fmt.Println( ( (a % 4) == 0 ) && ( (a % 100) != 0) )
    */

    /*
    fmt.Print("int dispari e compreso tra 0 e 100 ([0, 100]) : ")
    fmt.Scan(&a)
    fmt.Println( ( (a % 2) != 0) && ( (a >= 0) && (a <= 100)) )
    */

    var b float64
    fmt.Print("float64 'vicino' a 10.0 (entro 10^-6) (10e-6) : ")
    fmt.Scan(&b)
    fmt.Println(math.Abs(b-10.0) < 10e-6)}