package main

import "fmt"

func main() {
	var m, q float64
	var px, py float64
	var y float64

	fmt.Print("inserisci m e q:")
	fmt.Scan(&m, &q)

	fmt.Print("inserisci x e y:")
	fmt.Scan(&px, &py)

	y = m*px + q

	if py == y {
		fmt.Println("il punto appartiene alla retta")
	}
	if py > y {
		fmt.Println("il punto si trova sopra la retta")
	}
	if py < y {
		fmt.Println("il punto si trova sotto la retta")
	}
}
