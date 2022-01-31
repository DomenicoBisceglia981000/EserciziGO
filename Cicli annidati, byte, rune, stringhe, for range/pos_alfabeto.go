package main

import "fmt"

func main() {
	var c byte
	fmt.Println("Scrivi una sequenza di caratteri terminante con . :")
	for {
		fmt.Scanf("%d", &c)
		fmt.Print(string(c))
		break
	}
}
