package main

import "fmt"

func main() {

	var (
		a, b int8 = 30, 100     // overflow perchè troppo grande da rapprensentare con 8 bit
	)

	somma := a + b

	fmt.Println("La somma di", a, "e", b, "è", somma)

}
