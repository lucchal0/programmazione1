package main

import "fmt"

func main() {

	var s1 string = "Ma il cielo è sempre più blu!"

	var s2 []rune = []rune(s1)

	fmt.Println(len(s1), len(s2))
}
