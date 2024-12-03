package main

import (
	"fmt"
	"math"
)

func distanza(A, B struct{x float64; y float64}) (distanza_euclidea float64) {

	distanza_euclidea = math.Sqrt((A.x-B.x)*(A.x-B.x) + (A.y-B.y)*(A.y-B.y))

	return

}

func main() {

	var A struct {
		x float64
		y float64
	}
	var B struct {
		x float64
		y float64
	}

	A.x = 0.0
	B.x = 0.0
	A.y = 1.0
	B.y = 1.0

	fmt.Println("distanza tra due punti:", distanza(A,B))

}
