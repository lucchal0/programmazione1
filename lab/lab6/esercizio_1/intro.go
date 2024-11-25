package main

import "fmt"

func main() {
	var x rune = 'a'
	var y rune = 97
	var z rune = '\u0061'

	fmt.Println(x, y, z)
	fmt.Println()
	
	fmt.Println(string(x), string(y), string(z))

	fmt.Println()

	var beta1 rune = 946
	var beta2 rune = '\u0382'
	fmt.Println(string(beta1), string(beta2))

}
