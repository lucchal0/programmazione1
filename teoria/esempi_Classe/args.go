package main

import (
	"fmt"
	"os"
)

func main() {
	//var Args []string // slice di stringhe
	var s string

	//Args = make([]string,10)
	//Args[0]= "Ciao"
	//s= Args[0]

	for i := 0; i <len(os.Args); i++ {
		s = os.Args[i]
		fmt.Println(s)

	}
}
