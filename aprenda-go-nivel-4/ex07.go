package main

import "fmt"

func main()  {
	
	x := [][]string{
		[]string{"Robert", "Araújo Santos", "Football"},
		[]string{"Jonas", "Souza", "Corrida"},
		[]string{"Carlos", "Andrade", "Natação"},
	}

	for _, v := range x {
		fmt.Println(v)
	}
	
	fmt.Println("\n\n")

	for _, v := range x {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}
	}


	
}