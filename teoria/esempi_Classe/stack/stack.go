package stack

type Stack struct{
	pila []int 
	// eventuali altri elementi utili
}

func inserisci (s Stack, v int)(t Stack){
	t.pila = append(s.pila,v)

	return 
}

func vuota(s Stack) bool{
	return len(s.pila) == 0
}

func rimuovi(s Stack)(v int, s Stack){

	if vuota(s){
		
	}else{
		v = s.pila[len(s)-1]
		t.pila= s.pila[:len(s)-1]
	}
	return 

}