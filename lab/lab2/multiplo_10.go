package main

import "fmt"

func main() {
	var n int

	fmt.Print("inserisci un numero:")
	fmt.Scan(&n)

	if n%10 == 0 && n%10 != 0 {
		fmt.Println(n, " multiplo di 10")
	} else {
		fmt.Println(n, " non è multiplo di 10")
	}
}
