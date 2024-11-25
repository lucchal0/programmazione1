package main

import "fmt"

func main() {

	//var Asso_cuori int = '\U0001F0B1'
	var Asso_cuori rune = '\U0001F0B1'

	for x := Asso_cuori; x <= Asso_cuori+9; x++ {
		fmt.Println(string(x),"- codice formato in base 10:",x)
	}

}
