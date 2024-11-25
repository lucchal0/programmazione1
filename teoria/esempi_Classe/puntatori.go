package main

import "fmt"

func stampa_array(T [10]int) {
	for _, v := range T {
		fmt.Println(v)
	}
	T[0] = 45
}

func stampa_slice(T []int) {
	for _, v := range T {
		fmt.Println(v)
	}
	T[0] = 45
}

func main() {
	//var x int
	var S []int
	var A [10]int

	S = make([]int, 20)
	
	stampa_array(A)
	stampa_slice(S)
	
	fmt.Println("posizione 0 array:",A[0], "\n posizione 0 slice:",S[0])
}
