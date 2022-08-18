package main

import "fmt"

func main()  {
	x := 4
	// Função anonima
	func (x int)  {
		fmt.Println(x + 1)
	}(x)

}