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
		sobrenome: "Castro",
		sabores: []string{"chocolate", "avelã"},
	}

	x := map[string][]string{
        pessoa1.sobrenome : pessoa1.sabores,
		pessoa2.sobrenome : pessoa2.sabores,
	}

	for i, v := range x {
			fmt.Printf("%v\n" , i)
			for _, t := range v {
				fmt.Printf("\t%v\n" , t)
			}
	}


}