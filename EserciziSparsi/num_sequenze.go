package main
import "fmt"
func main(){
	var n, nPrec ,i int
	seq := false
	nPrec = 3
	for{
		fmt.Println("Inserisci 1 o 0 : ")
		fmt.Scan(&n)
		if n == 2{
			break
		}
		if seq == false && n == 1{
			seq = true
		}else if seq == false && n == 0{
			fmt.Print("LA SEQUENZA DEVE COMINCIARE E TERMINARE CON 1")
		}
		if seq == true && nPrec != n{
			if seq == true && n == 0{
				i ++
				
			}
		}
		nPrec = n
	}
	fmt.Print(i)
}