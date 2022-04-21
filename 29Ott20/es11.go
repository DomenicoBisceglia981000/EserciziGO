package main
import ("fmt"
	 "math/rand"
	 "time")
func main(){
	const k = 20
	var n, ris int
	rand.Seed(time.Now().UnixNano())
	for{
		n = rand.Intn(21)
		fmt.Println(n)
		ris = ris + n
		if n == k{
			break
		}
	}
	fmt.Println(ris)
}