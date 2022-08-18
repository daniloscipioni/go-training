package main

import "fmt"

type pessoa struct{
	nome string
	sobrenome string
	idade int
}

func (p pessoa) displayName() {
	fmt.Println("Nome Completo:", p.nome, p.sobrenome, "Idade:", p.idade)
}

func main()  {

	developer := pessoa{
		"Danilo Eduardo",
		"Scipioni",
		32,
	}

	developer.displayName()
}