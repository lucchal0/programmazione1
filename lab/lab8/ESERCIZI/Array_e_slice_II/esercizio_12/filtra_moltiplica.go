package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	nums:=[]int{}
	
	for _, v:=range os.Args[1:]{
		if n, err := strconv.Atoi(v); err== nil{
		nums = append(nums,n)
		}
	}

	if len(nums) >0 {
		fmt.Printf("il risultatato della moltiplicazione tra due numeri interi è:")

		ris:= nums[0]

		for i:=1; i<len(nums); i++{
			ris*=nums[i]
		}

		fmt.Printf("%d.\n",ris)

	}else{
		fmt.Println("non esistono numeri interi")
	}


}