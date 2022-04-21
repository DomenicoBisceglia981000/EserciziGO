package main
import "fmt"
import "math/rand"
import "time"
func main(){
	const k = 10

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < k; i++{
		fmt.Println(rand.Intn(11))
	}
}