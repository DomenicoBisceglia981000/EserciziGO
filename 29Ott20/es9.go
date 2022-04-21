package main
import "fmt"
import "math/rand"
import "time"
func main(){
	const TARGET = 50
	var n, ris int

	rand.Seed(time.Now().UnixNano())
	for ris < TARGET{
		n = rand.Intn(11)
		if ris + n <= TARGET{
			ris = ris + n
			fmt.Println(ris)
		}
	}
}