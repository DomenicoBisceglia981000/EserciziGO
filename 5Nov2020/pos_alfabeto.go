package main
import "fmt"
func main(){
	var ch byte
	fmt.Println("Inserisci una parola")
	for {
		fmt.Scanf("%c", &ch)
		if ch == '.'{
			break
		}
		if ch >= 'a' && ch <= 'z'{
			fmt.Println(string(ch), "è minuscola")
			fmt.Println(string(ch), "è nella posizione", ch - 'a')
		}
	}

	fmt.Println("bye")
}