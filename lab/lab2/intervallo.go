package main

import "fmt"

func main() {
	var voto int
	fmt.Print("VOTO = ")
	fmt.Scan(&voto)

	if voto <= 100 && voto >= 90 {
		fmt.Println("Ottimo")
	} else if voto >= 80 && voto < 90 {
		fmt.Println("Ottimo")
	} else if voto >= 70 && voto < 80 {
		fmt.Println("Distinto")
	} else if voto >= 60 && voto < 70 {
		fmt.Println("Sufficente")
	} else if voto < 60 && voto >= 0 {
		fmt.Println("Insufficente")
	} else if voto < 0 || voto > 100 {
		fmt.Println("Errore")
	}
}
