package main

import (
	"fmt"
	"math"
)

// type Punto = struct { // -> find replace
// 	x         float64
// 	y         float64
// 	etichetta string
// }

type Punto struct { // -> nuovo nome di tipo
	x         float64
	y         float64
	etichetta string
}

type celsius int32
type kelvin int32

func distanza(A, B Punto) (distanza_euclidea float64) {
	distanza_euclidea = math.Sqrt((A.x-B.x)*(A.x-B.x) + (A.y-B.y)*(A.y-B.y))
	return

}
func main() {

	var A Punto
	var B struct { // -> nuovo nome di tipo
		x         float64
		y         float64
		etichetta string
	}

	A.x = 0.0
	B.x = 0.0
	A.y = 1.0
	B.y = 1.0

	// var c celsius
	// var k kelvin

	B = A

	fmt.Println("distanza tra due punti:", distanza(A, B))

}
