package main
import "fmt"
func main(){
	const c = 10
	var max, cont, num int
	cont = 1
	for i:= 0; i < c; i++{
		fmt.Print("Inserire numero : ")
		fmt.Scan(&num)
		if num > max {
			max = num
			cont = 1
		}else if num == max{
			cont ++
		}



	}
	fmt.Print("Il numero max è : ", max, " e viene ripetuto ", cont, " volte")
}