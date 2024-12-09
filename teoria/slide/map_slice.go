package main

import "fmt"

func main() {
	var s, t []int
	s = make([]int, 6, 8) // lunghezzza variabile da 6 a 8
	s[0] = 10
	s[1] = 20
	s[2] = 30
	s[3] = 40
	s[4] = 50
	s[5] = 60
	t = s

	fmt.Println("s:", s)
	fmt.Println("t:", t)

	t = make([]int, 3, 5)
	fmt.Println("New t:", t)
	t = s[1:3:5] // da 1 a 3(escluso) e capacità 5-1=4
	fmt.Println("lunghezza :", len(t), "capacità:", cap(t), "t=s[1:3:5]->", t)
	t[0] = 70
	fmt.Println("lunghezza :", len(t), "capacità:", cap(t), "t[0] = 70->", t)
	Q:= append(t,80)
	fmt.Println("lunghezza :", len(Q), "capacità:", cap(Q), "Q:= append(t,80)->", Q)
	Q= append(t,91,92,93 )
	fmt.Println("lunghezza :", len(Q), "capacità:", cap(Q), "Q:= append(t,80)->", Q)

	fmt.Println("----------------------------------------------------------------------------------")

	mappe()
}

func mappe()bool{
	var m map[string]int
    m = make(map[string]int)

    m["alberto"] = 33
	//m["Alice"]=10
    fmt.Println("m:", m)

    var ok bool 
    e := m["alberto"]
    fmt.Println("Value of 'alberto':", e)

    e, ok = m["alberto"]
    fmt.Println("Value of 'alberto':", e, "Exists:", ok)
	e, ok = m["Alice"]
    fmt.Println("Value of 'Alice':", e, "Exists:", ok)
	return true

}

