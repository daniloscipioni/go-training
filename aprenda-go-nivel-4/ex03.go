package main

import "fmt"

func main()  {

	firstSlice := []int{1,2,3,4,5,6,7,8,9,105}

	segundoSlice := firstSlice[:3]
	terceiroSlice := firstSlice[4:]
	quartoSlice := firstSlice[1:7]
	quintoSlice := firstSlice[2:9]
	desafioSlice := firstSlice[2:len(firstSlice) - 1]

	fmt.Printf("Valor Slice Original: %v  \n", firstSlice)
	fmt.Printf("segundoSlice: %v  \n", segundoSlice)
	fmt.Printf("terceiroSlice: %v  \n", terceiroSlice)
	fmt.Printf("quartoSlice: %v  \n", quartoSlice)
	fmt.Printf("quintoSlice: %v  \n", quintoSlice)
	fmt.Printf("desafioSlice: %v  \n", desafioSlice)
	
	
}