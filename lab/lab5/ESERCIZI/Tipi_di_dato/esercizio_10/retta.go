package main

import (
	"fmt"
	"math"
	"math/rand"
)

const EPSILON = 1.0e-9

func XMaggioreDiY(x, y float64)bool {
	return (x - y) > EPSILON
}

func XUgualeAY(x, y float64) bool {
	return math.Abs(x-y) <= EPSILON
}

func XMinoreDiY(x, y float64) bool {
	return (x - y) < -EPSILON
}

func main() {
	var s int64

	fmt.Print("inserisci s:")
	fmt.Scan(&s)

	var m, q float64
	fmt.Print("Inserisci m e q: ")
	fmt.Scan(&m, &q)

	rand.Seed(s)

	for i := 0; i < 10; i++ {
		px := rand.Float64() * 10.0
		py := rand.Float64() * 10.0

		Linea:= m * px + q

		if XUgualeAY(py, Linea) {
			fmt.Printf("Punto (%f, %f) - Il punto sta sulla retta\n", px, py)
		} else if XMaggioreDiY(py, Linea) {
			fmt.Printf("Punto (%f, %f) - Il punto sta sopra la retta\n", px, py)
		} else if XMinoreDiY(py, Linea) {
			fmt.Printf("Punto (%f, %f) - Il punto sta sotto la retta\n", px, py)
		}
	}
}
