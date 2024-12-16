// stampa da 10 a 1

package main

import "fmt"
import "strconv"
import "os"


func conta_permutazioni(N int) int{
	if N == 1 {
		return 
	}
	var k int 
	k = N-1
	v:= conta_permutazioni(k)
	res:= k * v
	return res

}

func conta_permutazioni_ciclo(N int) int {
	var res int 
	res = 1 

	for i:= 0; i <N; i++{
		res = res * i
	}
	return res 
}

func member_it(s []int, v int)bool{
	for _,q := range s{
		if q == v{
			return true
		}else{
			return false
		}
	}
}



func countdown(N int) {

	var k int 
	if N < 1 {
		fmt.Println(N)
		return
	}
	fmt.Println("+++",N)
	k = N-1
	countdown(k)
	fmt.Println("---",N)
	return
}

func main() {
	var N int

	N, _ = strconv.Atoi(os.Args[1])

	countdown(N)
	conta_permutazioni(N)

}
