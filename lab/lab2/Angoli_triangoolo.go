package main

import "fmt"

func main() {
	var a, b int
	var c int

	fmt.Print("inserire ampiezza di due angoli: ")
	fmt.Scan(&a, &b)

	c = 180 - (a + b)

	if a+b+c == 180 && c > 0 {
		fmt.Println("Ampiezza secondo angolo =", c, "°")
	} else {
		fmt.Println("i due angoli non appartengono al triangolo")
	}

}
