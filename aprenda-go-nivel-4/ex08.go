package main

import "fmt"

func main()  {

	x := map[string][]string{
			"danilo_scipioni": []string{"jogar bola","assistir tv","teste"}, 
			"danilo_eduardo": []string{"andar com o cachorro","Ler","teste"}}


	for k, v := range x {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h )
		}
	}		
}