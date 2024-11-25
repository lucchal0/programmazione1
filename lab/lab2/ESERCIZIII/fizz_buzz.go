package main

import "fmt"

func main() {
	var n int

	fmt.Print("inserisci un numero intero:")
	fmt.Scan(&n)

	if n > 0 {
		if n % 3==0 {
			fmt.Print("Fizz ")
		}
		if n % 5==0 {
			fmt.Println("Buzz")
		}
	}

}
