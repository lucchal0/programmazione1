package main

import "fmt"

func main() {
	var Distanza float64

	fmt.Println("inserire la distanza in km=")
	fmt.Scan(&Distanza)

	fmt.Print("distanza in (mi) =", Distanza*0.621371)
}
