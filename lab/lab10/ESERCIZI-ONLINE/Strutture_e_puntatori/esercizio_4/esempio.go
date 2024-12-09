package main

import "fmt"

func main() {
	var ptr *int
	ptr = new(int)
	*ptr = 50
	fmt.Println(*ptr)
}
