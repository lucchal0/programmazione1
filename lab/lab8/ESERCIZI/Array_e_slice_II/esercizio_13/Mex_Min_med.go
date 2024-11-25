package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	nums := []int{}

	for _, v := range os.Args[:] {
		if n, err := strconv.Atoi(v); err == nil {
			nums = append(nums, n)
		}
	}

	fmt.Printf("Minimo: %d.\n", Minimo(nums[:]))
	fmt.Printf("Massimo: %d.\n", Massimo(nums[:]))
	fmt.Printf("Media: %f.\n", Media(nums[:]))

}

func Minimo(sl []int) int {
	var min int
	min = sl[0]

	for i := 0; i < len(sl); i++ {
		if min >= sl[i] {
			min = sl[i]
		}
	}
	return min
}

func Massimo(sl []int) int {
	var max int
	max = sl[0]

	for i := 0; i < len(sl); i++ {
		if max <= sl[i] {
			max = sl[i]
		}
	}
	return max
}

func Media(sl []int) float64 {
	var med float64
	var c float64

	for i := 0; i < len(sl); i++ {
		med = float64(sl[i]) + med
		c++
	}

	return med / c
}
