package main

import "fmt"
import "strconv"
import "math/rand"
import "os"
import "time"

func mescola(N int) (S []int) {

	S = make([]int, N)

	for i := 0; i < len(S); i++ {
		S[i] = i
	}

	ora_corrente := time.Now().UnixNano()
	rand.Seed(ora_corrente)

	for i := len(S) - 1; i > 0; i-- {
		// scegliere un indice tra 0 e i
		// rand.Intn(....)
		j := rand.Intn(i + 1)
		//v := S[i]
		//S[i] = S[j]
		//S[j] = v
		S[i], S[j] = S[j], S[i]
	}

	return S

}

func leggiInput(argomenti []string) (s, c int) {

	s, _ = strconv.Atoi(argomenti[1])
	c, _ = strconv.Atoi(argomenti[2])

	return

}

// stampaMazzo:
// Input: [0 4 5 3 1 2], 2 , 3
// Output: -
// Effetto collaterale: stampa del tipo (0, 0) (1, 1) (1, 2) (1, 0) (0, 1) (0, 2)

func stampaMazzo(mazzo []int, s int, c int) {
	for _, card := range mazzo {
		s := card / c
		v := card % c
		fmt.Printf("(%d, %d) ", s, v)
	}
	fmt.Println()
}

func main() {
	var semi, carte, N int
	var mazzo []int

	semi, carte = leggiInput(os.Args)
	fmt.Println(semi, carte)

	N = semi * carte

	mazzo = mescola(N)
	fmt.Println(mazzo)

	stampaMazzo(mazzo, semi, carte)

}
