package main

import "fmt"


func main()  {

	slice:=[]int{1,2,3,4,900}

	// Do indice 0 ao indice 3 - Exclui o indice 3
	//fatia := slice[0:3]
	// Todos os valores do slice
	//fatia := slice[:]
	// Do indice 1 ao final do slice
	fatia := slice[1:len(slice)]

	fmt.Println(fatia)
}