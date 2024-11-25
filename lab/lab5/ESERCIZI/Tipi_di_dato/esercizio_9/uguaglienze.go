package main

import "fmt"
import "math"

func main() {
	var x float64

	fmt.Scan(&x)

	r := math.Sqrt(x)

	if r*r == x {
		fmt.Println("uguali")
	} else {
		fmt.Println("divesri")
	}

}
