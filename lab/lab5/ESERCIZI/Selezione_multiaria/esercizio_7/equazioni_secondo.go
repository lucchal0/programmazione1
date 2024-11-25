package main

import "fmt"
import "math"

func main() {

	var a, b, c float64

	fmt.Scan(&a, &b, &c)

	delta := b*b - 4*a*c

	switch {
		case delta > 0:

			z1 := -b + math.Sqrt(delta)/(2*a)
			z2 := -b - math.Sqrt(delta)/(2*a)

			fmt.Println("risultati radici relai:", z1, z2)

		case delta == 0:
			fmt.Println("delta = 0")

			z1 := -b / (2 * a)

			fmt.Println("risultati radici relai:", z1)

		case delta < 0:
			fmt.Println("nessuna soluzione nei numeri in R")

			z1 := (complex(-b, 0) + complex(0, math.Sqrt(-delta))) / complex(2*a, 0)
			z2 := (complex(-b, 0) - complex(0, math.Sqrt(-delta))) / complex(2*a, 0)

			fmt.Println("risultati radici:", z1, z2)
	}

}
