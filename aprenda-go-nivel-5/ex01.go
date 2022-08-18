package main

import "fmt"

type pessoa struct {
	nome string;
	sobrenome string;
	sabores []string
}

func main()  {

	pessoa1:= pessoa{
		nome: "Danilo",
		sobrenome: "Scipioni",
		sabores: []string{"creme", "morango"},
	}

	pessoa2:= pessoa{
		nome: "Doraci",
		sobrenome: "de Castro",
		sabores: []string{"chocolate", "avelã"},
	}

	fmt.Printf("-- %v\n",pessoa1.nome)
	for _, v := range pessoa1.sabores {
		
		fmt.Printf("\t%v\n", v)
	}

	fmt.Printf("-- %v\n",pessoa2.nome)
	for _, v := range pessoa2.sabores {
		
		fmt.Printf("\t%v\n", v)
	}


}