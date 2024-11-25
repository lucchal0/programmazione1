package main

import "fmt"

func main() {
	/*
	   g -> giorno
	   m -> mese
	   a -> anno
	*/
	var g, m, a int

	//Printf utilizzo per stampare
	fmt.Printf("ciao a tutti\n")

	g = 30
	m = 10
	a = 24

	/**
	  nel Printf posso utilizzare %d per stampare decimali
	*/

	fmt.Printf("Oggi è ancora %d/%d/%d ottobre\n", g, m, a)

	fmt.Printf("Il giorno, in base dieci è: %d\n", g)

	//definiamo un numero decimale
	g = 0b11110

	// Nello stesso modo che utilizzimao %d possimao utilizzare:
	// %b -> stampa binario (0,1)
	// %x -> stampa in esadecimale (0,1,2,3,4,5,6,7,8,9,a,b,c,d,e,f)

	fmt.Printf("Il giorno, in base dieci è: %d; il giorno in base 2, è %b; in base 16: %x\n", g, g, g)

	/*
	   	OPERAZIONI bit a bit

	   AND

	   	1100 in binario
	   	1010 in binario

	   result := a & b     1000 in binario

	   -----------------------------------------------------------------------------

	   OR

	   	1100 in binario
	   	1010 in binario

	   result := a | b     1110 in binario (14 in decimale)

	   -----------------------------------------------------------------------------

	   NOT

	   	1100 in binario

	   result := ^a      0011 in binario (-13 in decimale per la rappresentazione a complemento a due)

	   -----------------------------------------------------------------------------

	   	operatore di SHIFT

	   	0010 << 1 -> 0100    -> l'ultimo bit va perso e in quella vota ci metto lo 0
	   	0010 >> 1 -> 0001
	*/
	var val32 int32
	var i int

	val32 = 0b00000000000000000000000000000001

	for i = 1; i <= 35; i++ {

		fmt.Printf("bit %d %d %.32b\n", i, val32, val32)
		val32 = val32 << 1
	}

}
