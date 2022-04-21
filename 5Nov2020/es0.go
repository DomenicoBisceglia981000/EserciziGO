package main
import "fmt"

func main(){
	var ch rune

	fmt.Print("Inserisci un carattere : ")
	fmt.Scanf("%c", &ch)
	//Punto1
	fmt.Println(string(ch), "è una runa")
	//Punto2
	fmt.Println(string(ch -1), string(ch), string(ch + 1))
	//Punto3
	if (ch >= 'a' && ch <= 'l') || (ch >= 'A' && ch <= 'L'){
		fmt.Println("Compreso fra A --> L")
		
	}else{
		fmt.Println("Altro")
	}
	//Punto4
	var s string
	fmt.Println("Inserisci una parola")
	fmt.Scan(&s)
	/*
	for i := 0; i < len(s);  i++{
		fmt.Println(string(s[i]))
	}
	*/
	for i,  r := range s{
		fmt.Println(string(r))
		_ = i
	}
}