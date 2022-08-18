package main

import "fmt"


type pessoa struct{
	nome string
	idade int
}

// *pessoa tipo ponteiro = com * é o tipo ponteiro
func mudeMe(p *pessoa)  {
	fmt.Println("Valor", p)
	(*p).nome = "Rodrigo"
}

func main()  {
	pessoa := pessoa{
		"Danilo",
		32,
	}

	fmt.Printf("ValorMain %T", &pessoa)
	//Ponteiro para pessoa = &pessoa
	mudeMe(&pessoa)
	fmt.Println(pessoa)
}