package main

import "fmt"

func structSemplice() {
	var p struct {
		x         float64
		y         float64
		etichetta string
	}

	var p1 struct {
		x         float64
		y         float64
		etichetta string
	}

	p.x = 23
	p.y = 55
	p.etichetta = "23+55"

	fmt.Println("struct p: x:", p.x, " y:", p.y, " etichetta:", p.etichetta)
	fmt.Println("struct p1: x:", p1.x, " y:", p1.y, " etichetta:", p1.etichetta)
	fmt.Println()

	p.x = p.x * 2
	fmt.Println("struct p: x*2:", p.x, " y:", p.y)

	println("somma:", somma(p))

	println("\n p1 = p")

	p1 = p
	fmt.Println("struct p1: x:", p1.x, " y:", p1.y, " etichetta:", p1.etichetta)
}

func somma(p struct {
	x         float64
	y         float64
	etichetta string
}) int {
	somma := p.x + p.y
	return int(somma)
}

//------------------------------------------------------------------------------------------------------

type cordinate struct {
	x float64
	y float64
}

type SliceofInt = []int

func Typefunction() float64 {
	var a cordinate

	var B cordinate

	a.x = 1
	a.y = 4
	B.x = 3

	fmt.Println("type cordinate: x:", a.x, " y:", a.y)
	fmt.Println("type cordinate: x:", B.x, " y:", B.y)

	somma2 := somma2(a, B)

	fmt.Println("somma delle x: ", somma2)

	return somma2

}

func somma2(a, B cordinate) float64 {
	var somma float64
	somma = a.x + B.x
	return somma

}

// ----------------------------------------

type Animale interface {
	verso() string
}

type cane struct{}

func (c cane) verso() string {
	return "bau"
}

type gatto struct{}

func (g gatto) verso() string {
	return "miau"
}

func manipolazione_Stringhe() {
	var animali []Animale = []Animale{cane{}, gatto{}}
	for _, animal := range animali {
		fmt.Println(animal.verso())
	}
}

func main() {

	Typefunction()

	fmt.Println()
	fmt.Println("-------------------------------------------------")
	fmt.Println()

	structSemplice()
	fmt.Println()
	manipolazione_Stringhe()

}
