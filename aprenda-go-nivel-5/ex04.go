package main

import "fmt"

func main()  {
	
x := struct {
	nome string
	sabor string
	pessoa map[int]string
	slice []int
}{	
	 nome: "Testando Struct anonima",
	 sabor: "testando sabor",
	 pessoa: map[int]string { 1: "danilo", 2: "eduardo" },
	 slice: []int {1,2,3,4,5},
}
	
fmt.Println(x)

}