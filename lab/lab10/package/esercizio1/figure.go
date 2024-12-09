package main

import "./rettangolo"
import "fmt"

func main(){


	var a = rettangolo.NewRettangolo(10.0, 20.0)

	// base, altezza := 10.0, 20.0
	// r:= rettangolo.NewRettangolo(base, altezza)
	fmt.Println(rettangolo.Area(*a))
}
