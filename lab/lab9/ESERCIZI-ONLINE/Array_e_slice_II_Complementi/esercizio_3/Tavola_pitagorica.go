package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])

	fmt.Println(n)

	tavola := creaTavolaPita(n)

	fmt.Printf("%#v\n", tavola)
}

func creaTavolaPita(n int) [][]int {
	var tavola [][]int

	tavola = make([][]int, n)
	for i := 0; i < n; i++ {
		tavola[i] = make([]int, n)
	}

	for i:=0; i<n; i++{
		for j:=0; j<n; j++{
			tavola[j][j]= 
		}
	}

	return tavola
}
