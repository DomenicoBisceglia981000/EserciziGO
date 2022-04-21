package main
import "fmt"
type data struct{
	g, m, a int
}
func main(){
	var d data
	fmt.Println(d)
}

func ultFeb(anno int) int{
	var s data
	return data{28, 2, anno}
}